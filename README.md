# dynamic-compilation

## Install dependencies

Install golang migrate

```sh
$ go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Run dependencies

Start a local `mysql` db:

```sh
$ docker compose up -d
```

Export enviroment variables:

```sh
$ export MYSQL_DSN="dev:dev@tcp(127.0.0.1:3306)/demo?parseTime=true"
```

Run migrations:

```sh
$ migrate -database mysql://${MYSQL_DSN} -path ./migrations up
```

Run demo:

```sh
$ go run .
```
