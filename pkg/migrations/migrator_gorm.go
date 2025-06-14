package migrations

import (
	"context"
	"errors"
	"fmt"
	"github.com/axidex/api-example/pkg/tables"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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

func (m *MigratorGorm) Migrate(schemaName, ownerName string) error {
	models := []interface{}{
		&tables.User{},
	}

	err := m.CreateSchema(schemaName, ownerName)
	if err != nil {
		return err
	}

	err = m.db.Migrator().AutoMigrate(models...)
	if err != nil {
		return err
	}

	for _, model := range models {
		tableModel, ok := model.(schema.Tabler)
		if !ok {
			return errors.New("interface of table is not a tabler")
		}
		tableName := tableModel.TableName()
		comment := getTableComment(model)
		m.db.Exec("COMMENT ON TABLE " + tableName + " IS '" + comment + "';")
	}

	return nil
}

func getTableComment(model interface{}) string {
	switch model.(type) {
	case *tables.User:
		return "Table for storing user information."

	default:
		return ""
	}
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
