package main

import (
	"flag"
	"github.com/acger/consumer-svc/route"
	"github.com/tal-tech/go-zero/core/conf"
	"log"
	"os"
	"os/signal"
	_ "regexp"

	cluster "github.com/bsm/sarama-cluster"
)

var configFile = flag.String("f", "etc/cs.yaml", "the config file")
var logger *log.Logger

func main() {
	flag.Parse()

	var c route.Config
	conf.MustLoad(*configFile, &c)
	ctx := route.NewCtx(c)

	file, err := os.OpenFile(c.Log.Path + "error.log", os.O_APPEND|os.O_CREATE, 666)
	if err != nil {
		log.Fatalln("fail to create error.log file!")
	}
	defer file.Close()
	logger = log.New(file, "", log.LstdFlags|log.Lshortfile)

	// init (custom) config, enable errors and notifications
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	// init consumer
	brokers := c.Queue.Host
	topics := c.Queue.Topic
	consumer, err := cluster.NewConsumer(brokers, c.Queue.Group, topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			logger.Println(err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Println(ntf)
		}
	}()

	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				route.Route(msg, ctx)
			}
		case <-signals:
			return
		}
	}
}
