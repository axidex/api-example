package db

import (
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

func WithTelemetry(db *gorm.DB, tp trace.TracerProvider) (*gorm.DB, error) {
	if err := db.Use(
		tracing.NewPlugin(
			tracing.WithTracerProvider(tp),
		),
	); err != nil {
		return nil, fmt.Errorf("can't create tracing for gorm: %w", err)
	}

	return db, nil
}
