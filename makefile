all: protoc_build go_build

protoc_build: protoc_integration protoc_light

protoc_integration:
	protoc --go_out=./pkg/apis/core/integration --go_opt=paths=source_relative --go-grpc_out=./pkg/apis/core/integration --go-grpc_opt=paths=source_relative -I=./pkg/apis/core/integration ./pkg/apis/core/integration/integration.proto

protoc_light:
	protoc --go_out=./pkg/apis/core/light --go_opt=paths=source_relative --go-grpc_out=./pkg/apis/core/light --go-grpc_opt=paths=source_relative -I=./pkg/apis/core/light -I=./pkg/apis/core/integration ./pkg/apis/core/light/light.proto

go_build:
	go build

go_test:
	go clean -testcache
	go test -count=1 --test.v ./...
	go clean -testcache