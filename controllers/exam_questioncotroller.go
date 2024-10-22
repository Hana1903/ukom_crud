package controllers

import (
	"net/http"
	"time"
	"ukom_crud/config"
	"ukom_crud/models"
	"github.com/gin-gonic/gin"
)

// Create ExamQuestion
func CreateExamQuestion(c *gin.Context) {
	var input struct {
		ExamID     int64  `json:"exam_id"`
		QuestionID  int64  `json:"question_id"`
		UserAnswer string `json:"user_answer"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create new ExamQuestion
	examQuestion := models.ExamQuestion{
		ExamID:     input.ExamID,
		QuestionID: input.QuestionID,
		UserAnswer: input.UserAnswer,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Save examQuestion to the database
	if err := config.DB.Create(&examQuestion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, examQuestion)
}

// Get all exam questions
func GetExamQuestions(c *gin.Context) {
	var examQuestions []models.ExamQuestion
	if err := config.DB.Find(&examQuestions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, examQuestions)
}

// Get a single exam question by ID
func GetExamQuestionByID(c *gin.Context) {
	var examQuestion models.ExamQuestion
	if err := config.DB.Where("id = ?", c.Param("id")).First(&examQuestion).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ExamQuestion not found"})
		return
	}
	c.JSON(http.StatusOK, examQuestion)
}

// Update exam question by ID
func UpdateExamQuestion(c *gin.Context) {
	var examQuestion models.ExamQuestion
	if err := config.DB.Where("id = ?", c.Param("id")).First(&examQuestion).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ExamQuestion not found"})
		return
	}

	var input struct {
		ExamID     int64  `json:"exam_id"`
		QuestionID  int64  `json:"question_id"`
		UserAnswer string `json:"user_answer"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update examQuestion details
	examQuestion.ExamID = input.ExamID
	examQuestion.QuestionID = input.QuestionID
	examQuestion.UserAnswer = input.UserAnswer
	examQuestion.UpdatedAt = time.Now()

	if err := config.DB.Save(&examQuestion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, examQuestion)
}

// Delete exam question by ID
func DeleteExamQuestion(c *gin.Context) {
	var examQuestion models.ExamQuestion
	if err := config.DB.Where("id = ?", c.Param("id")).First(&examQuestion).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ExamQuestion not found"})
		return
	}

	if err := config.DB.Delete(&examQuestion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ExamQuestion deleted successfully"})
}
