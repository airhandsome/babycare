package controllers

import (
    "net/http"
    "time"

    "github.com/babycare/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type GrowthController struct {
    db *gorm.DB
}

func NewGrowthController(db *gorm.DB) *GrowthController {
    return &GrowthController{db: db}
}

// CreateRecord 创建成长记录
func (gc *GrowthController) CreateRecord(c *gin.Context) {
    var record models.GrowthRecord
    if err := c.ShouldBindJSON(&record); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    record.UserID = 1 // TODO: 从JWT获取用户ID
    record.CreatedAt = time.Now()
    record.UpdatedAt = time.Now()

    if err := gc.db.Create(&record).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, record)
}

// GetRecords 获取成长记录
func (gc *GrowthController) GetRecords(c *gin.Context) {
    var records []models.GrowthRecord
    userID := uint(1) // TODO: 从JWT获取用户ID
    recordType := c.Query("type")
    startDate := c.Query("start_date")
    endDate := c.Query("end_date")

    query := gc.db.Model(&models.GrowthRecord{}).
        Where("user_id = ?", userID)

    if recordType != "" {
        query = query.Where("type = ?", recordType)
    }

    if startDate != "" {
        query = query.Where("date >= ?", startDate)
    }

    if endDate != "" {
        query = query.Where("date <= ?", endDate)
    }

    if err := query.Order("date desc").Find(&records).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, records)
}

// UpdateMilestone 更新里程碑状态
func (gc *GrowthController) UpdateMilestone(c *gin.Context) {
    var milestone models.Milestone
    if err := c.ShouldBindJSON(&milestone); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := gc.db.Model(&models.Milestone{}).
        Where("id = ?", c.Param("id")).
        Updates(map[string]interface{}{
            "title":        milestone.Title,
            "category":     milestone.Category,
            "completed":    milestone.Completed,
            "completed_at": milestone.CompletedAt,
            "updated_at":   time.Now(),
        }).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 返回更新后的里程碑
    if err := gc.db.First(&milestone, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, milestone)
}

// CreateMilestone 创建里程碑
func (gc *GrowthController) CreateMilestone(c *gin.Context) {
    var milestone models.Milestone
    if err := c.ShouldBindJSON(&milestone); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // TODO: 从JWT获取用户ID
    milestone.UserID = 1
    milestone.CreatedAt = time.Now()
    milestone.UpdatedAt = time.Now()

    if err := gc.db.Create(&milestone).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, milestone)
}

// GetMilestones 获取里程碑列表
func (gc *GrowthController) GetMilestones(c *gin.Context) {
    var milestones []models.Milestone
    userID := uint(1) // TODO: 从JWT获取用户ID

    query := gc.db.Model(&models.Milestone{}).
        Where("user_id = ?", userID)

    if err := query.Order("category, created_at").Find(&milestones).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, milestones)
} 