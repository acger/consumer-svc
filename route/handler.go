package route

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/acger/consumer-svc/model"
)

type Routers struct {
}

func (router *Routers) PairChat(msg *sarama.ConsumerMessage, ctx *Ctx, args ...interface{}) bool {
	m := model.Chat{}
	json.Unmarshal(msg.Value, &m)
	ctx.ChatDataBase.Create(&m)

	return true
}
