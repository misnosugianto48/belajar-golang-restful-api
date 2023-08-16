### Package

1. github.com/go-playground/validator/v10
2. github.com/julienschmidt/httprouter
3. github.com/sirupsen/logrus
4. github.com/stretchr/testify
5. https://github.com/golang-migrate/migrate

### Create Migrate

migrate create -ext sql -dir db/migrations create_table_name

### Run Migrate

#### Up

migrate -database "mysql://db_user@tcp(db_host:db_port)/db_name" -path db/migrations up

#### Down

migrate -database "mysql://db_user@tcp(db_host:db_port)/db_name" -path db/migrations down

// use number to migrate each version //

### Version

migrate -database "mysql://db_user@tcp(db_host:db_port)db_name" -path db/migrations force migrate_version

// force migrate to change schema_migrations in database where new version has a dirty state //
