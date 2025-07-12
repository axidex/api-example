package migrations

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var ErrNotTablerInterface = errors.New("interface of table is not a tabler")

type MigratorGorm struct {
	db *gorm.DB
}

func CreateMigrator(ctx context.Context, db *gorm.DB) Migrator {
	return &MigratorGorm{
		db: db.WithContext(ctx),
	}
}

func (m *MigratorGorm) Migrate(models []any, schemaName, ownerName string) error {
	err := m.CreateSchema(schemaName, ownerName)
	if err != nil {
		return err
	}

	err = m.db.Migrator().AutoMigrate(models...)
	if err != nil {
		return err
	}

	return nil
}

func (m *MigratorGorm) CreateTable(table interface{}) error {

	if !m.db.Migrator().HasTable(table) {
		err := m.db.Migrator().CreateTable(table)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MigratorGorm) CreateSchema(schemaName, ownerName string) error {
	createSchemaSQL := fmt.Sprintf(`
        CREATE SCHEMA IF NOT EXISTS %s;
    `, schemaName) // SQL INJECTION NOT EXPLOITABLE, BECAUSE IT'S CONFIG VARIABLE
	res := m.db.Exec(createSchemaSQL)
	if res.Error != nil {
		return res.Error
	}

	alterSchemaOwnerSQL := fmt.Sprintf(`
        ALTER SCHEMA %s OWNER TO %s;
    `, schemaName, ownerName) // SQL INJECTION NOT EXPLOITABLE, BECAUSE IT'S CONFIG VARIABLES
	res = m.db.Exec(alterSchemaOwnerSQL)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
