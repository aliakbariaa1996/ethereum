build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

build:
	docker-compose build ethereum

run:
	docker-compose up ethereum

test:
	go test -v ./...

migrate:
	migrate -path ./db/migration -database migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/authmarch?sslmode=disable" -verbose up

swag:
	swag init -g cmd/main.go