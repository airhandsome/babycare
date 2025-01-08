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

type ForumController struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewForumController(db *gorm.DB, rdb *redis.Client) *ForumController {
	return &ForumController{db: db, rdb: rdb}
}

// CreatePost 创建帖子
func (fc *ForumController) CreatePost(c *gin.Context) {
	var post models.ForumPost
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT获取用户ID
	post.UserID = 1
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	if err := fc.db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// GetPosts 获取帖子列表
func (fc *ForumController) GetPosts(c *gin.Context) {
	var posts []models.ForumPost
	var total int64
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	category := c.Query("category")

	offset := (page - 1) * pageSize
	query := fc.db.Preload("User")

	if category != "" && category != "all" {
		query = query.Where("category = ?", category)
	}

	// 获取总数
	if err := query.Model(&models.ForumPost{}).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := query.Order("created_at desc").Offset(offset).Limit(pageSize).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
		"total": total,
		"page": page,
		"page_size": pageSize,
	})
}

// LikePost 点赞帖子
func (fc *ForumController) LikePost(c *gin.Context) {
	id := c.Param("id")
	var post models.ForumPost

	if err := fc.db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if err := fc.db.Model(&post).Update("likes", gorm.Expr("likes + ?", 1)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Liked successfully"})
}

// CreateComment 创建评论
func (fc *ForumController) CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// 设置评论属性
	comment.PostID = uint(postID)
	comment.UserID = 1 // TODO: 从JWT获取用户ID
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	// 开启事务
	tx := fc.db.Begin()
	if err := tx.Create(&comment).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 更新帖子评论数
	if err := tx.Model(&models.ForumPost{}).Where("id = ?", postID).
		Update("comments", gorm.Expr("comments + ?", 1)).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回完整的评论信息（包括用户信息）
	if err := fc.db.Preload("User").First(&comment, comment.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// GetComments 获取评论列表
func (fc *ForumController) GetComments(c *gin.Context) {
	postID := c.Param("id")
	var comments []models.Comment

	if err := fc.db.Where("post_id = ?", postID).
		Preload("User").
		Order("created_at desc").
		Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comments,
		"total": len(comments),
	})
}

// LikeComment 点赞评论
func (fc *ForumController) LikeComment(c *gin.Context) {
	commentID := c.Param("id")
	var comment models.Comment

	// 开启事务
	tx := fc.db.Begin()

	// 查找评论
	if err := tx.First(&comment, commentID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// 更新点赞数
	if err := tx.Model(&comment).Update("likes", gorm.Expr("likes + ?", 1)).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回更新后的评论信息
	if err := fc.db.Preload("User").First(&comment, commentID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}
