package route

import (
	"github.com/acger/consumer-svc/tool"
	"gorm.io/gorm"
)

type Ctx struct {
	ChatDataBase *gorm.DB
}

func NewCtx(config Config) *Ctx {
	return &Ctx{
		ChatDataBase: tool.NewMysql(config.ChatDataSource),
	}
}
