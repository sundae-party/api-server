# Build

## protoc build integration

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. -go-grpc_opt=paths=source_relative --proto_path=. integration.proto
```

## protoc build light

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. -go-grpc_opt=paths=source_relative --proto_path=. --proto_path=/go/src/api_server/pkg/apis/core/integration light.proto
```

## Test

Create new mongo for the tests:

```docker
docker run --rm -it --name mongo-test -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=pwd -e MONGO_INITDB_DATABASE=sundae -e MONGO_INITDB_USERNAME=sundae -e MONGO_INITDB_PASSWORD=pass mongo
```

```bash
go test -test.v ./...
```
