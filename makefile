all: protoc_build go_build

protoc_build: protoc_integration protoc_light protoc_binary_sensor protoc_sensor protoc_sun

protoc_integration:
	protoc --go_out=./pkg/apis/core/types --go_opt=paths=source_relative --go-grpc_out=./pkg/apis/core/types --go-grpc_opt=paths=source_relative -I=./pkg/apis/core/integration ./pkg/apis/core/integration/integration.proto

protoc_light:
	protoc --go_out=./pkg/apis/core/types --go_opt=paths=source_relative --go-grpc_out=./pkg/apis/core/types --go-grpc_opt=paths=source_relative -I=./pkg/apis/core/light -I=./pkg/apis/core/integration ./pkg/apis/core/light/light.proto

protoc_binary_sensor:
	protoc --go_out=./pkg/apis/core/types --go_opt=paths=source_relative --go-grpc_out=./pkg/apis/core/types --go-grpc_opt=paths=source_relative -I=./pkg/apis/core/binary_sensor -I=./pkg/apis/core/integration ./pkg/apis/core/binary_sensor/binary_sensor.proto

protoc_sensor:
	protoc --go_out=./pkg/apis/core/types --go_opt=paths=source_relative --go-grpc_out=./pkg/apis/core/types --go-grpc_opt=paths=source_relative -I=./pkg/apis/core/sensor -I=./pkg/apis/core/integration ./pkg/apis/core/sensor/sensor.proto

protoc_sun:
	protoc --go_out=./pkg/apis/core/types --go_opt=paths=source_relative --go-grpc_out=./pkg/apis/core/types --go-grpc_opt=paths=source_relative -I=./pkg/apis/core/sun -I=./pkg/apis/core/integration ./pkg/apis/core/sun/sun.proto

go_build:
	go build

go_test:
	go clean -testcache
	go test -count=1 --test.v ./...
	go clean -testcache