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

	for {
		resp := client.Watch(context.Background(), "logger")
		for wresp := range resp {
			for _, ev := range wresp.Events {
				fmt.Printf("%s %q :%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}

}
