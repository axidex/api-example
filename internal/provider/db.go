package provider

import (
	"context"
	"github.com/axidex/api-example/pkg/db"
)

func (p *Provider) initDatabase(_ context.Context) error {
	psqlEngine, err := db.NewPostgresConnection(p.cfg.Database, p.logger)
	if err != nil {
		return err
	}

	p.dependencies.db = psqlEngine

	return nil
}
