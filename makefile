all: protoc_build go_build

protoc_build: protoc_integration protoc_entity

protoc_integration:
	protoc --go_out=./pkg/apis/core/types --go_opt=paths=source_relative --go-grpc_out=./pkg/apis/core/types --go-grpc_opt=paths=source_relative -I=./pkg/apis/core/integration ./pkg/apis/core/integration/integration.proto

protoc_entity:
	protoc --go_out=./pkg/apis/core/types --go_opt=paths=source_relative --go-grpc_out=./pkg/apis/core/types --go-grpc_opt=paths=source_relative -I=./pkg/apis/core/entity -I=./pkg/apis/core/integration ./pkg/apis/core/entity/entity.proto

go_build:
	go build

go_test:
	go clean -testcache
	go test -count=1 --test.v ./...
	go clean -testcache