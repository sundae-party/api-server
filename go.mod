module github.com/sundae-party/api-server

go 1.15

//replace google.golang.org/grpc => google.golang.org/grpc v1.27.0

//replace go.etcd.io/etcd => github.com/fredczj/etcd v3.4.6+incompatible

require (
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.2
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	go.mongodb.org/mongo-driver v1.4.4
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)
