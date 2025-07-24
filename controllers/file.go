package controllers

import (
	"fmt"
	"go-file-service/config"
	"go-file-service/models"
	"go-file-service/storage"
	"go-file-service/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	Storage *storage.LocalStorage
}

func (fc *FileController) Upload(c *gin.Context) {
	// 1. 获取上传文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "获取文件失败: " + err.Error(),
		})
		return
	}

	// 2. 打开文件流
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "打开文件失败: " + err.Error(),
		})
		return
	}
	defer src.Close()

	// 3. 读取文件内容
	fileBytes := make([]byte, file.Size)
	if _, err := src.Read(fileBytes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "读取文件内容失败: " + err.Error(),
		})
		return
	}

	// 4. 存储到本地
	filePath, err := fc.Storage.Save(file.Filename, fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "文件保存失败: " + err.Error(),
		})
		return
	}

	// 5. 准备数据库记录
	fileRecord := models.File{
		ID:   utils.GenerateUUID(),
		Name: file.Filename,
		Path: filepath.Base(filePath), // 只存储文件名，不包含路径
		Size: file.Size,
	}

	// 6. 保存到数据库（使用 config.DB）
	if err := config.DB.Create(&fileRecord).Error; err != nil {
		// 数据库操作失败时，删除已存储的文件
		_ = fc.Storage.Delete(filepath.Base(filePath))

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "数据库记录创建失败: " + err.Error(),
			"details": fmt.Sprintf(
				"尝试写入的记录: ID=%s, Name=%s, Path=%s, Size=%d",
				fileRecord.ID,
				fileRecord.Name,
				fileRecord.Path,
				fileRecord.Size,
			),
		})
		return
	}

	// 7. 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"id":   fileRecord.ID,
		"name": fileRecord.Name,
		"size": fileRecord.Size,
		"path": fileRecord.Path,
	})
}
