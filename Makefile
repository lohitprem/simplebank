test:
	go test -v -cover ./...

.PHONY: createdb deletedb migrateup migratedown sqlc test
