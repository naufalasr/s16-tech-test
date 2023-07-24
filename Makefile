.PHONY: generate
generate:
	@protoc --go_out=./generated --go_opt=paths=source_relative \
    --go-grpc_out=./generated --go-grpc_opt=paths=source_relative \
    ./*.proto

.PHONY: run
run:
	@go run cmd/main.go

.PHONY: test
test:
	@go test ./services --cover