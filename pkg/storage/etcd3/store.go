package etcd3

import (
	"context"
	"fmt"

	//"go.etcd.io/etcd/clientv3"
	"sundae-party/api-server/pkg/apis/core/types"
)

type EtcdStore struct {
	cli string
}

func NewStore() *EtcdStore {
	// c, err := clientv3.New(clientv3.Config{
	// 	Endpoints:   []string{"http://172.17.0.4:2379"},
	// 	DialTimeout: 2 * time.Second,
	// })
	// if err == context.DeadlineExceeded {
	// 	// handle errors
	// }

	// return Store{
	// 	Client: c,
	// }

	return &EtcdStore{}
}

// TODO ETCD functions implementing Store interface
func PutIntegration(ctx context.Context, newIntegration *types.Integration) (*types.Integration, error) {
	return &types.Integration{}, nil
}
func GetIntegration(ctx context.Context, name string) (*types.Integration, error) {
	return &types.Integration{}, nil
}
func DeleteIntegration(ctx context.Context, deleteIntegration *types.Integration) (string, error) {
	return fmt.Sprintf("%s deleted.", "TODO"), nil
}
