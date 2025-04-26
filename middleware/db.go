package middleware

import (
	"github.com/hurricane5250/go-project/log"
	"github.com/hurricane5250/go-project/model"
)

// MakeMigrations 迁移数据库表
func MakeMigrations() {
	var err error
	err = model.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&model.User{},
		)
	if err != nil {
		log.Error("[ error ] 生成数据库表失败: %v", err)
		return
	}
	log.Info("[ success ] 生成数据库表成功")
}
