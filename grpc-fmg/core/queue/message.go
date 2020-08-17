package queue

import (
	"grpc-demo/utils"
	"grpc-demo/utils/log"
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

var Message *nsq.Producer

//type NSQHandler struct {
//	name string
//}
//
//func (this *NSQHandler) HandleMessage(msg *nsq.Message) error {
//	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body), this.name)
//	return nil
//}

// 初始化消息队列
func InitMessageQueue() {
	url := fmt.Sprintf("%s:%d", utils.GlobalConfig.Nsq.Host, utils.GlobalConfig.Nsq.Port)
Conn:
	producer, err := nsq.NewProducer(url, nsq.NewConfig())
	if err != nil {
		logUtils.Println("消息队列连接失败， 五秒后尝试重连:", err)
		time.Sleep(time.Second * 5)
		goto Conn
	}
	Message = producer


	//go func() {
	//	config:=nsq.NewConfig()
	//	config.MaxInFlight=9
	//
	//
	//	consumer, err := nsq.NewConsumer("test", "struggle", config)
	//	if nil != err {
	//		fmt.Println("err", err)
	//		return
	//	}
	//	consumer.AddHandler(&NSQHandler{name:"xiotng"})
	//	err = consumer.ConnectToNSQD(url)
	//	if nil != err {
	//		fmt.Println("err", err)
	//		return
	//	}
	//	select {}
	//}()
}