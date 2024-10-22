package controllers

import (
	"net/http"
	"regexp"
	"time"
	"ukom_crud/config"
	"ukom_crud/models"
	"github.com/gin-gonic/gin"
)

// Create a new user
func CreateUser(c *gin.Context) {
	var input struct {
		Name                   string `json:"Name"`
		Email                  string `json:"Email"`
		Password               string `json:"Password"`
		DateOfBirth            string `json:"DateOfBirth"`
		Gender                 string `json:"Gender"`
		PhoneNumber            string `json:"PhoneNumber"`
		EducationalInstitution string `json:"EducationalInstitution"`
		Profession             string `json:"Profession"`
		Address                string `json:"Address"`
		Province               string `json:"Province"`
		City                   string `json:"City"`
	}

	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate phone number (only digits allowed)
	phoneNumberPattern := `^[0-9]+$`
	matched, err := regexp.MatchString(phoneNumberPattern, input.PhoneNumber)
	if err != nil || !matched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number. Only digits are allowed."})
		return
	}

	// Parse DateOfBirth from string to time.Time
	layout := "2006-01-02"
	dob, err := time.Parse(layout, input.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}

	// Create a new User struct
	user := models.User{
		Name:                   input.Name,
		Email:                  input.Email,
		Password:               input.Password,
		DateOfBirth:            models.DateOfBirth(dob),
		Gender:                 input.Gender,
		PhoneNumber:            input.PhoneNumber,
		EducationalInstitution: input.EducationalInstitution,
		Profession:             input.Profession,
		Address:                input.Address,
		Province:               input.Province,
		City:                   input.City,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}

	// Save the user to the database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created user
	c.JSON(http.StatusOK, user)
}

// // Get all users
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// Get a user by ID
func GetUserByID(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update a user by ID
func UpdateUser(c *gin.Context) {
	var user models.User
	// Check if the user exists
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind JSON input
	var input struct {
		Name                   string `json:"Name"`
		Email                  string `json:"Email"`
		Password               string `json:"Password"`
		DateOfBirth            string `json:"DateOfBirth"`
		Gender                 string `json:"Gender"`
		PhoneNumber            string `json:"PhoneNumber"`
		EducationalInstitution string `json:"EducationalInstitution"`
		Profession             string `json:"Profession"`
		Address                string `json:"Address"`
		Province               string `json:"Province"`
		City                   string `json:"City"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate phone number (only digits allowed)
	phoneNumberPattern := `^[0-9]+$`
	matched, err := regexp.MatchString(phoneNumberPattern, input.PhoneNumber)
	if err != nil || !matched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number. Only digits are allowed."})
		return
	}

	// Parse DateOfBirth from string to time.Time
	layout := "2006-01-02"
	dob, err := time.Parse(layout, input.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}

	// Update the user fields
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	user.DateOfBirth = models.DateOfBirth(dob)
	user.Gender = input.Gender
	user.PhoneNumber = input.PhoneNumber
	user.EducationalInstitution = input.EducationalInstitution
	user.Profession = input.Profession
	user.Address = input.Address
	user.Province = input.Province
	user.City = input.City
	user.UpdatedAt = time.Now()

	// Save the updated user
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the updated user
	c.JSON(http.StatusOK, user)
}

// Delete a user by ID
func DeleteUser(c *gin.Context) {
	var user models.User
	// Check if the user exists
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Delete the user
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
