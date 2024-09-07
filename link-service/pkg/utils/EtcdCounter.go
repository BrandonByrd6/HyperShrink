package utils

import (
	"go.etcd.io/etcd/clientv3"
)

type EtcdCounter struct {
	cli *clientv3.KV
}

func NewEtcdCounter(cli *clientv3.KV) *EtcdCounter {
	return &EtcdCounter{}
}
