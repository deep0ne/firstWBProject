postgres:
	docker run --name wbpostgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=wbpass -d postgres

createdb:
	docker exec -it wbpostgres createdb --username=root --owner=root wborders

dropdb:
	docker exec -it wbpostgres dropdb wborders

migrateup:
	migrate -path db/migration -database "postgresql://root:wbpass@localhost:5432/wborders?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:wbpass@localhost:5432/wborders?sslmode=disable" -verbose down

nats:
	docker run --network host -p 4222:4222 -d nats -js

server:
	go run main.go

testdb:
	go test ./db/

.PHONY: postgres createdb dropdb migrateup migratedown nats testdb