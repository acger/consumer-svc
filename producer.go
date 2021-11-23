package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"time"
)

var Address = []string{"127.0.0.1:9092"}

func main()  {
	var msg string
	for  {
		fmt.Scanln(&msg)
		SaramaProducer(msg, "chat")
	}
}

func SaramaProducer(message ,topic string) {
	syncProducer(Address, message,topic)
}

//同步消息模式
func syncProducer(address []string, message string,topic string) {
	//指定配置
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewSyncProducer(address, config)

	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
	defer p.Close()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}
	part, offset, err := p.SendMessage(msg)
	if err != nil {
		log.Printf("send message(%s) err=%s \n", message, err)
	} else {
		fmt.Fprintf(os.Stdout, message+"发送成功，partition=%d, offset=%d \n", part, offset)
	}

}