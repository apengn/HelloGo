package etcd

import (
	"time"
	etctClient "github.com/coreos/etcd/clientv3"
	"os"
	"github.com/astaxie/beego/logs"
	"HelloGo/logagent/tail"
	"context"
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
	return
}

func PutToEtcd(msg tail.TextMsg) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	cancel()
	_, err := client.Put(ctx, msg.Topic, msg.Msg)
	if os.IsExist(err) {
		logs.Error("put etcd fail", msg.Topic, msg.Msg)
	}
}
