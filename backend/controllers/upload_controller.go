package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UploadController struct {
	uploadPath string
}

func NewUploadController(uploadPath string) *UploadController {
	return &UploadController{uploadPath: uploadPath}
}

// UploadImages 处理图片上传
func (uc *UploadController) UploadImages(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := form.File["images"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No images uploaded"})
		return
	}

	urls := make([]string, 0, len(files))
	for _, file := range files {
		// 生成唯一文件名
		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
		
		// 保存文件
		dst := filepath.Join(uc.uploadPath, filename)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 生成访问URL
		urls = append(urls, fmt.Sprintf("/uploads/%s", filename))
	}

	c.JSON(http.StatusOK, gin.H{
		"urls": urls,
	})
} 