package controllers

import (
	"net/http"
	"time"
	"ukom_crud/config"
	"ukom_crud/models"
	"github.com/gin-gonic/gin"
)

// Fungsi untuk parsing string waktu menjadi time.Time
func parseTime(input string) (time.Time, error) {
	// Ubah format waktu dari "YYYY-MM-DD HH:MM:SS" ke "YYYY-MM-DDTHH:MM:SS"
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// Create Exam
func CreateExam(c *gin.Context) {
	var input struct {
		OrderID   int     `json:"order_id"`
		PackageID int     `json:"package_id"`
		UserID    int     `json:"user_id"`
		Name      string  `json:"name"`
		StartedAt string  `json:"started_at"`  
		EndedAt   string  `json:"ended_at"`    
		Score     float64 `json:"score"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parsing waktu input menjadi time.Time
	startedAt, err := parseTime(input.StartedAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid started_at format"})
		return
	}

	endedAt, err := parseTime(input.EndedAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ended_at format"})
		return
	}

	// Create new Exam
	exam := models.Exam{
		OrderID:   input.OrderID,
		PackageID: input.PackageID,
		UserID:    input.UserID,
		Name:      input.Name,
		StartedAt: startedAt,
		EndedAt:   endedAt,
		Score:     input.Score,
	}

	// Save exam to the database
	if err := config.DB.Create(&exam).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exam)
}

// Get all exams
func GetExams(c *gin.Context) {
	var exams []models.Exam
	if err := config.DB.Find(&exams).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, exams)
}

// Get a single exam by ID
func GetExamByID(c *gin.Context) {
	var exam models.Exam
	if err := config.DB.Where("id = ?", c.Param("id")).First(&exam).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exam not found"})
		return
	}
	c.JSON(http.StatusOK, exam)
}

// Update exam by ID
func UpdateExam(c *gin.Context) {
	var exam models.Exam
	if err := config.DB.Where("id = ?", c.Param("id")).First(&exam).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exam not found"})
		return
	}

	var input struct {
		OrderID   int     `json:"order_id"`
		PackageID int     `json:"package_id"`
		UserID    int     `json:"user_id"`
		Name      string  `json:"name"`
		StartedAt string  `json:"started_at"`  
		EndedAt   string  `json:"ended_at"`    
		Score     float64 `json:"score"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parsing waktu input menjadi time.Time
	startedAt, err := parseTime(input.StartedAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid started_at format"})
		return
	}

	endedAt, err := parseTime(input.EndedAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ended_at format"})
		return
	}

	// Update exam details
	exam.OrderID = input.OrderID
	exam.PackageID = input.PackageID
	exam.UserID = input.UserID
	exam.Name = input.Name
	exam.StartedAt = startedAt
	exam.EndedAt = endedAt
	exam.Score = input.Score

	if err := config.DB.Save(&exam).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exam)
}

// Delete exam by ID
func DeleteExam(c *gin.Context) {
	var exam models.Exam
	if err := config.DB.Where("id = ?", c.Param("id")).First(&exam).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exam not found"})
		return
	}

	if err := config.DB.Delete(&exam).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exam deleted successfully"})
}
