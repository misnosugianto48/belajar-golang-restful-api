### Package

1. https://github.com/go-playground/validator/v10
2. https://github.com/julienschmidt/httprouter
3. https://github.com/sirupsen/logrus
4. https://github.com/stretchr/testify
5. https://github.com/golang-migrate/migrate
6. https://github.com/go-sql-driver/mysql

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
