serve:
	go run main.go serve

generate-ids:
	go run main.go generate:ids

create-migration:
	goose -dir build/migrations create $(file) sql

run-migrations:
	goose -dir build/migrations mysql "${MYSQL_DSN}" up

generate-proto:
	protoc ./api/proto/*.proto --proto_path=api/proto --go-grpc_out=. --go_out=.
