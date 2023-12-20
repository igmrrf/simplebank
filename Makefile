postgres:
	docker run --name some-postgress -e POSTGRES_PASSWORD=secretpassword -e POSTGRES_USER=root -d -p 5432:5432 postgres

droppostgres:
	docker container stop some-postgress && docker container rm some-postgress

createdb:
	docker exec -it some-postgress createdb --username=root --owner=root simple_bank 

dropdb:
	docker exec -it some-postgress dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres dropdb createdb migrateup migratedown sqlc test
