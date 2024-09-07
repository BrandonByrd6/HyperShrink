package connections

import (
	"fmt"
	"log"
	"time"

	"github.com/brandonbyrd6/link-service/pkg/config"
	"go.etcd.io/etcd/clientv3"
)

func Init() clientv3.KV {
	cfg := config.GetConfig()
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://10.50.0.21:2379"},
		DialTimeout: time.Duration(cfg.Etcd.DialTimeOut) * time.Second,
	})
	if err != nil {
		fmt.Println("etcd")
		log.Fatalln(err)
	}

	kv := clientv3.NewKV(client)
	return kv
}
