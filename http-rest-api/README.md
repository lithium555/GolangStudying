./apiserver --help

Usage of ./apiserver:
  -config-path string
        path to config file (default "configs/apiserver.toml")

migrate create -ext sql -dir http-rest-api/migrations -seq create_users
migrate -path migrations -database "postgres://localhost/restapi_dev?sslmode=disable" up

createdb restapi_test
migrate -path migrations -database "postgres://localhost/restapi_test?sslmode=disable" up

psql -d restapi_dev
