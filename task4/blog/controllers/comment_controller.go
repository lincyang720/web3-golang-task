package controllers

import (
	"net/http"
	"strconv"
	"blog/database"
	"blog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateComment 创建评论
func CreateComment(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 设置评论作者和文章ID
	comment.UserID = userID.(uint)

	// 检查文章是否存在
	var post models.Post
	if err := database.DB.First(&post, comment.PostID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// 创建评论
	if err := database.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	// 加载关联的用户和文章信息
	database.DB.Preload("User").Preload("Post").First(&comment, comment.ID)

	c.JSON(http.StatusCreated, comment)
}

// GetCommentsByPost 获取某篇文章的所有评论
func GetCommentsByPost(c *gin.Context) {
	postID := c.Param("post_id")
	
	// 检查文章ID是否有效
	if _, err := strconv.Atoi(postID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// 检查文章是否存在
	var post models.Post
	if err := database.DB.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var comments []models.Comment
	
	// 获取文章的所有评论并预加载用户信息
	if err := database.DB.Where("post_id = ?", postID).Preload("User").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}