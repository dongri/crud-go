# crud-go

## infra
```
$ cd infra
$ docker network create crud-go_network
$ docker-compose up --build -d
$ docker-compose logs -f
```

## server
```
$ cd server
$ go install github.com/pressly/goose/v3/cmd/goose@latest

$ goose -dir db postgres "host=localhost port=5432 user=postgres dbname=crud-go sslmode=disable" up

$ docker-compose up --build -d
$ docker-compose logs -f
```

## front
```
$ cd front
$ yarn install
$ yarn start
```

http://localhost:3000

## Demo
https://crud-go.onrender.com/
