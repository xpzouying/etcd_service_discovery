package main

import (
	"context"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 3 * time.Second,
	})
	panicIfErr(err)
	defer cli.Close()

	// minimum lease TTL is 5-second
	resp, err := cli.Grant(context.TODO(), 5)
	panicIfErr(err)

	// after 5 seconds, the key 'foo' will be removed
	_, err = cli.Put(context.TODO(), "service/worker1", "here is worker1", clientv3.WithLease(resp.ID))
	panicIfErr(err)
}
