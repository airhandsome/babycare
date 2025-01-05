package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/babycare/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type PostController struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewPostController(db *gorm.DB, rdb *redis.Client) *PostController {
	return &PostController{db: db, rdb: rdb}
}

// CreatePost 创建文章
func (pc *PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 模拟用户ID (实际应从JWT token中获取)
	post.UserID = 1
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	if err := pc.db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// GetPosts 获取文章列表
func (pc *PostController) GetPosts(c *gin.Context) {
	var posts []models.Post
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	category := c.Query("category")

	offset := (page - 1) * pageSize
	query := pc.db.Preload("User")

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// GetPost 获取单个文章
func (pc *PostController) GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := pc.db.Preload("User").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// UpdatePost 更新文章
func (pc *PostController) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := pc.db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	var updateData models.Post
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.Title = updateData.Title
	post.Content = updateData.Content
	post.UpdatedAt = time.Now()

	if err := pc.db.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

// DeletePost 删除文章
func (pc *PostController) DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := pc.db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if err := pc.db.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

// IncrementViews 增加文章阅读量
func (pc *PostController) IncrementViews(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := pc.db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if err := pc.db.Model(&post).Update("views", gorm.Expr("views + ?", 1)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Views updated successfully"})
}

// GetPostsByCategory 获取指定分类的文章
func (pc *PostController) GetPostsByCategory(c *gin.Context) {
	category := c.Query("category")
	var posts []models.Post

	query := pc.db.Preload("User")
	if category != "" {
		query = query.Where("category = ?", category)
	}

	if err := query.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
} 