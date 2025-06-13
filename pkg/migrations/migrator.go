package migrations

type Migrator interface {
	Migrate(schemaName, ownerName string) error
}
