package route

import (
	"github.com/Shopify/sarama"
	"github.com/acger/consumer-svc/tool"
	"reflect"
	"strings"
)

type FuncCollection map[string]reflect.Value

func CallFunc(tableName string, args ...interface{}) (result []reflect.Value, err error) {
	var router Routers
	FuncMap := make(FuncCollection, 0)
	rf := reflect.ValueOf(&router)
	rft := rf.Type()
	funcNum := rf.NumMethod()
	for i := 0; i < funcNum; i++ {
		mName := rft.Method(i).Name
		FuncMap[mName] = rf.Method(i)
	}
	parameter := make([]reflect.Value, len(args))
	for k, arg := range args {
		parameter[k] = reflect.ValueOf(arg)
	}
	result = FuncMap[tableName].Call(parameter)
	return
}

func Route(msg *sarama.ConsumerMessage, ctx *Ctx) {
	topic := msg.Topic
	items := strings.Split(topic, "-")

	for i, item := range items {
		items[i] = tool.FirstToUpper(item)
	}

	CallFunc(strings.Join(items, ""), msg, ctx)
}
