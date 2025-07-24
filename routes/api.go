package routes // 包名必须与目录名一致

import (
	"go-file-service/controllers"
	"go-file-service/storage"

	"github.com/gin-gonic/gin"
)

func SetupFileRoutes(r *gin.Engine) {
	fileStorage := storage.NewLocalStorage("uploads")
	fileCtrl := controllers.FileController{Storage: fileStorage}

	api := r.Group("/api")
	{
		api.POST("/upload", fileCtrl.Upload)
	}
}
