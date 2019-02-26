package main

import (
	etctClient "github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
	"context"
)

func main() {
	client, err := etctClient.New(etctClient.Config{Endpoints: []string{
		"localhost:2379",
		"localhost:22379",
		"localhost:32379",
	},
		DialTimeout: 5 * time.Second})

	if err != nil {
		fmt.Println("connect etcd failed", err)
		return
	}
	fmt.Println("connect etcd succ", err)

	defer client.Close()
	var count int
	for {
		count++
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := client.Put(ctx, "nginx_log", "sample_value"+string(count))
		cancel()
		if err != nil {
			fmt.Println("put etcd kv failed", err)
			return
		}
		time.Sleep(time.Second)
		fmt.Println(resp.OpResponse(),"===")
	}
	//ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	//resp, err := client.Get(ctx, "logger")
	//
	//cancel()
	//if err != nil {
	//	fmt.Println("get etcd kv failed", err)
	//	return
	//}
	//
	//for k, v := range resp.Kvs {
	//	fmt.Printf("index %d  etcd key%s:%s\n", string(k), string(v.Key), string(v.Value))
	//}

}
