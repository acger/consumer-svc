package model

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Uid     uint64 `gorm:"index" json:"uid"`
	ToUid   uint64 `gorm:"index" json:"to_uid"`
	Message string `gorm:"size:3000" json:"message"`
	Status  bool   `json:"status"`
}

type ChatHistory struct {
	gorm.Model
	Uid   uint64 `gorm:"index"`
	ToUid uint64
}
