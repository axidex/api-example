package app

import (
	"context"
	"github.com/axidex/api-example/internal/config"
	"github.com/axidex/api-example/pkg/logger"
	"github.com/axidex/api-example/pkg/telemetry"
	"github.com/oklog/run"
	"syscall"
	"time"
)

type initFunc func(context.Context) error

type IApp interface {
	Run(ctx context.Context) error
}

type App struct {
	telemetry telemetry.Telemetry
	cfg       *config.Config
	logger    logger.Logger
	name      string
	debug     bool
}

func NewApp() IApp {
	return &App{}
}

func (a *App) Run(ctx context.Context) error {
	if err := a.init(ctx); err != nil {
		return err
	}
	defer a.stop()

	g := &run.Group{}
	//g.Add(a.handler.HandleServer(ctx))
	g.Add(run.SignalHandler(ctx, syscall.SIGINT, syscall.SIGTERM))

	if err := g.Run(); err != nil {
		//a.logger.Errorf("Error occured | %s", err.Error())
		return err
	}

	return nil
}

func (a *App) init(ctx context.Context) error {

	inits := []initFunc{
		a.initConfig,
		a.initName,
		a.initTelemetry,
		a.initLogger,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			return err
		}
	}

	go a.testMetrics(ctx)
	go a.testLogs(ctx)

	return nil
}

func (a *App) testMetrics(ctx context.Context) {
	mp := a.telemetry.GetMeterProvider()
	counter, _ := mp.Meter("test").Int64Counter("test.counter")

	for {
		counter.Add(ctx, 1)
		time.Sleep(10 * time.Second)
	}
}

func (a *App) testLogs(ctx context.Context) {
	for {
		a.testLog(ctx)
	}
}

func (a *App) testLog(ctx context.Context) {
	ctx, span := a.telemetry.GetTracerProvider().Tracer("test").Start(ctx, "test.logs")
	defer span.End()
	a.logger.Info(ctx, "Test logs - %d", 10)
	a.qwe(ctx)

	time.Sleep(10 * time.Second)
}

func (a *App) qwe(ctx context.Context) {
	ctx, span := a.telemetry.GetTracerProvider().Tracer("internal").Start(ctx, "internal.qwe")
	defer span.End()

	a.logger.Info(ctx, "Internal test logs - %d", 10)
}

func (a *App) stop() {
	a.telemetry.Stop(context.Background())
}
