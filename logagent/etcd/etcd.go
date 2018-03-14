package etcd

import (
	"time"
	etctClient "github.com/coreos/etcd/clientv3"
	"os"
	"github.com/astaxie/beego/logs"
	"context"
	"fmt"
)

var (
	client *etctClient.Client
)

func InitEtcd(Endpoint []string) (err error) {
	client, err = etctClient.New(etctClient.Config{Endpoints: Endpoint,
		DialTimeout: 5 * time.Second})
	if os.IsExist(err) {
		logs.Error(err)
		return
	}
	logs.Info("etcd connect success")
	return
}

func PutToEtcd(topic,msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	cancel()
	res, err := client.Put(ctx, topic, msg)
	if os.IsExist(err) {
		panic(err)
		logs.Error("put etcd fail", topic, msg)
	}

	fmt.Printf("put etcd  topic==%v   msg:=====%v\n",topic,msg)
	fmt.Println(res.OpResponse(),"================")
}
