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

```bash
go test -test.v ./...
```
