migrateup:
	migrate -path database/migration -database "postgresql://admin:password@127.0.0.1:5435/db?sslmode=disable" -verbose up

migrateup-test:
	migrate -path database/migration -database "postgresql://admin:password@127.0.0.1:5435/db_test?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration -database "postgresql://admin:password@127.0.0.1:5435/db?sslmode=disable" -verbose down

migratedown-test:
	migrate -path database/migration -database "postgresql://admin:password@127.0.0.1:5435/db_test?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: migrateup migratedown sqlc test migrateup-test migratedown-test
