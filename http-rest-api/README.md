./apiserver --help

Usage of ./apiserver:
  -config-path string
        path to config file (default "configs/apiserver.toml")

migrate create -ext sql -dir http-rest-api/migrations -seq create_users
migrate -path migrations -database "postgres://localhost/restapi_dev?sslmode=disable" up

createdb restapi_test
migrate -path migrations -database "postgres://localhost/restapi_test?sslmode=disable" up

psql -d restapi_dev

http POST http://localhost:8080/users email=invalid
http POST http://localhost:8080/users email=user@example.org password=password

http POST http://localhost:8080/users email=user@example.org password=123


http http://localhost:8080/private/whoami
http -v --session=user POST http://localhost:8080/sessions email=user@example.org password=password
http -v  --session=user http://localhost:8080/private/whoami
http --session=user http://

http -v  --session=user http://localhost:8080/private/whoami "Origin: google.com"