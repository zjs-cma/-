package config

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// 导出全局DB变量（首字母大写）
var DB *gorm.DB

func InitDB() {
	dsn := "sqlserver://sa:188089@localhost:1433?database=go"
	var err error
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
}
