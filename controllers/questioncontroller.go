package controllers

import (
	"ukom_crud/config"
	"ukom_crud/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Create a new question
func CreateQuestion(c *gin.Context) {
	var input struct {
		IDPackage     int64  `json:"id_package" binding:"required"`
		Question      string `json:"question" binding:"required"`
		Answer        string `json:"answer"`
		CorrectAnswer string `json:"correct_answer"`
	}

	// Bind the JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question := models.Question{
		IDPackage:     input.IDPackage,
		Question:      input.Question,
		Answer:        input.Answer,
		CorrectAnswer: input.CorrectAnswer,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// Save the question to the database
	if err := config.DB.Create(&question).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, question)
}

// Get all questions
func GetQuestions(c *gin.Context) {
	var questions []models.Question
	config.DB.Find(&questions)
	c.JSON(http.StatusOK, questions)
}

// Get a question by ID
func GetQuestionByID(c *gin.Context) {
	var question models.Question
	if err := config.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}
	c.JSON(http.StatusOK, question)
}

// Update a question by ID
func UpdateQuestion(c *gin.Context) {
	var question models.Question
	if err := config.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	var input struct {
		IDPackage     int64  `json:"id_package" binding:"required"`
		Question      string `json:"question" binding:"required"`
		Answer        string `json:"answer"`
		CorrectAnswer string `json:"correct_answer"`
	}

	// Bind the JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	question.IDPackage = input.IDPackage
	question.Question = input.Question
	question.Answer = input.Answer
	question.CorrectAnswer = input.CorrectAnswer
	question.UpdatedAt = time.Now()

	// Save the updated question to the database
	config.DB.Save(&question)
	c.JSON(http.StatusOK, question)

	if err := config.DB.Save(&question).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, question)
}

// Delete a question by ID
func DeleteQuestion(c *gin.Context) {
	var question models.Question
	if err := config.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	config.DB.Delete(&question)
	c.JSON(http.StatusOK, gin.H{"message": "Question deleted successfully"})
}