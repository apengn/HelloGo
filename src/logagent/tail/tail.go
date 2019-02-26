package tail

import (
	"github.com/hpcloud/tail"
	"time"
	"github.com/astaxie/beego/logs"
	"HelloGo/logagent/etcd"
	"sync"
)

type CollectConf struct {
	LogPath string
	Topic   string
}

type TextMsg struct {
	Msg   string
	Topic string
}

type TailObj struct {
	tail     *tail.Tail
	conf     CollectConf
	status   int
	exitChan chan int
}

type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan  chan *TextMsg
	lock     sync.Mutex
}

var (
	tailObjMgr *TailObjMgr
)

func InitTail(conf []CollectConf, chanSize int) {

	if len(conf) == 0 {
		logs.Error("invalid config for log collect, conf:%v", conf)
	}
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}

	for _, v := range conf {
		createNewTask(v)
	}
}
func createNewTask(conf CollectConf) (err error) {

	tailObj := TailObj{
		conf: conf,
	}

	tails, err := tail.TailFile(conf.LogPath, tail.Config{
		ReOpen: true,
		Follow: true,
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		logs.Error("tail file err:", err)
		return

	}
	tailObj.tail = tails
	tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, &tailObj)
	for true {
		if !readFromTail(tailObj) {
			logs.Warn("tail file close reopen, filename:%s\n", tailObj.tail.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
	}

	return
}
func readFromTail(tailObj TailObj) (ok bool) {
	line, ok := <-tailObj.tail.Lines
	if !ok {
		return
	}
	msg := &TextMsg{
		Msg:   line.Text,
		Topic: tailObj.conf.Topic,
	}

	//tailObjMgr.msgChan <- msg
	//msg1 := *GetFromChan()
	go etcd.PutToEtcd(msg.Topic, msg.Msg)

	return ok
}

func GetFromChan() *TextMsg {

	return <-tailObjMgr.msgChan
}
