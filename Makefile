.PHONY:  up
up:
	docker run --name avito_db -p 8080:5432 -e POSTGRES_PASSWORD=dev -d postgres:15.1-alpine

down:
	docker rm -f avito_db
