module sundae-party/api-server

go 1.15

//replace google.golang.org/grpc => google.golang.org/grpc v1.27.0

//replace go.etcd.io/etcd => github.com/fredczj/etcd v3.4.6+incompatible

require (
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.0
	github.com/gorilla/mux v1.8.0
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/zap v1.16.0 // indirect
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)
