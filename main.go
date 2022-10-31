package main

import (
	"flag"
	"fmt"
	"github.com/aliyunmq/mq-http-go-sdk"
	_ "go.mongodb.org/mongo-driver/bson"
	"log"
	"strconv"
	"time"
)

func main1() {
	var name string
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "go语言", "帮助信息")
		_ = goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "name", "php语言", "帮助信息")
		_ = phpCmd.Parse(args[1:])
	}
	log.Printf("name:%s", name)
}

type Error  struct {
	Code int
	Res string
}

func(e Error) Error() string {
	return ""
}

func ops() error {
	var err *Error
	return err
}

func main() {
err := ops()
fmt.Println(err)




	// 设置HTTP协议客户端接入点，进入消息队列RocketMQ版控制台实例详情页面的接入点区域查看。
	endpoint := "http://1566301846632690.mqrest.cn-beijing.aliyuncs.com"
	// AccessKey ID阿里云身份验证，在阿里云RAM控制台创建。
	accessKey := "LTAI4Fzq5GKmmj5XH3xTqCwp"
	// AccessKey Secret阿里云身份验证，在阿里云RAM控制台创建。
	secretKey := "FxX9hQnZ6qzT6KeuihnKHm1WJG8hKT"
	// 消息所属的Topic，在消息队列RocketMQ版控制台创建。
	topic := "wm_study_center"
	// Topic所属的实例ID，在消息队列RocketMQ版控制台创建。
	// 若实例有命名空间，则实例ID必须传入；若实例无命名空间，则实例ID传入null空值或字符串空值。实例的命名空间可以在消息队列RocketMQ版控制台的实例详情页面查看。
	instanceId := "MQ_INST_1566301846632690_BcvtClKQ"

	client := mq_http_sdk.NewAliyunMQClient(endpoint, accessKey, secretKey, "")

	mqProducer := client.GetProducer(instanceId, topic)
	// 循环发送4条消息。
	for i := 0; i < 4; i++ {
		var msg mq_http_sdk.PublishMessageRequest

		msg = mq_http_sdk.PublishMessageRequest{
			MessageBody: "hello mq!",         //消息内容。
			MessageTag:  "",                  // 消息标签。
			Properties:  map[string]string{}, // 消息属性。
		}
		// 设置消息的Key。
		msg.MessageKey = "MessageKey"
		// 设置消息自定义属性。
		msg.Properties["a"] = strconv.Itoa(i)

		ret, err := mqProducer.PublishMessage(msg)

		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Printf("Publish ---->\n\tMessageId:%s, BodyMD5:%s, \n", ret.MessageId, ret.MessageBodyMD5)
		}
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	//err := cmd.Execute()
	//if err != nil {
	//	log.Fatalf("cmd.Execute err: %v", err)
	//}
}
