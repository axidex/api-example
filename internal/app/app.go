package app

import (
	"context"
	"github.com/axidex/ss-manager/pkg/telemetry"
	"github.com/oklog/run"
	"syscall"
)

type initFunc func(context.Context) error

type IApp interface {
	Run(ctx context.Context) error
}

type App struct {
	telemetry telemetry.Telemetry
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
		a.initTelemetry,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) stop() {
	a.telemetry.Stop(context.Background())
}
