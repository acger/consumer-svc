package tool

import (
	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewMysql(dataSource string) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dataSource,
		DefaultStringSize: 255,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		logx.Error("mysql connect fail")
		return nil
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
