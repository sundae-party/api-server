package etcd3

import (
	"fmt"
	"regexp"
	//"go.etcd.io/etcd/clientv3"
)

type Store struct {
	//Client *clientv3.Client
	Data map[string]string
}

func NewStore() *Store {
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

	data := map[string]string{
		"/zwave": "{Name: zLight}",
	}

	return &Store{Data: data}
}

func (s Store) Put(k string, d string) (string, error) {
	//s.Client.KV.Put(ctx context.Context, key string, val string, opts ...clientv3.OpOption)
	s.Data[k] = d
	return s.Data[k], nil
}

func (s Store) Get(k string) (string, error) {
	return s.Data[k], nil
}

func (s Store) GetByIntegration(i string) []string {
	var resp []string
	for k, v := range s.Data {
		regStr := fmt.Sprintf("^/%s/", i)
		var regLight = regexp.MustCompile(regStr)
		if regLight.MatchString(k) {
			resp = append(resp, v)
		}
	}
	return resp
}
