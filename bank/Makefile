createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb bank --username=root --owner=root simple_bank

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb sqlc test
