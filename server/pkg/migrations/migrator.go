package migrations

type Migrator interface {
	Migrate(models []any, schemaName, ownerName string) error
}
