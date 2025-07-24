package main

import (
	"fmt"
	"go-file-service/config"
	"go-file-service/models"
	"go-file-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库（直接调用，不赋值）
	config.InitDB()
	// 新增：检查DB是否有效
	if config.DB == nil {
		fmt.Println("config.DB 是 nil！")
	} else {
		// 尝试获取底层数据库连接，进一步验证
		sqlDB, err := config.DB.DB()
		if err != nil {
			fmt.Println("获取底层连接失败:", err)
		} else {
			fmt.Println("底层数据库连接:", sqlDB)
			// 测试连接是否能ping通
			if err := sqlDB.Ping(); err != nil {
				fmt.Println("数据库ping失败:", err)
			} else {
				fmt.Println("数据库连接正常！")
			}
		}
	}
	// 自动迁移模型（确保models包中有AutoMigrate函数）
	if err := models.AutoMigrate(config.DB); err != nil {
		panic("数据库迁移失败: " + err.Error())
	}

	// 创建Gin实例
	r := gin.Default()
	routes.SetupFileRoutes(r)
	r.Run(":8084")
}
