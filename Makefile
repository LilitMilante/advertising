.PHONY:  up down migrate-install migrate
up:
	docker run --name avito_db -p 8080:5432 -e POSTGRES_PASSWORD=dev -d postgres:15.1-alpine

down:
	docker rm -f avito_db

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate

migrate:
	migrate -path=migrations/ -database postgres://postgres:dev@localhost:8080/postgres?sslmode=disable up
