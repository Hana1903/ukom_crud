package controllers

import (
	"net/http"
	"regexp"
	"time"
	"ukom_crud/config"
	"ukom_crud/models"
	"github.com/gin-gonic/gin"
)

func UpdateUserProfile(c *gin.Context) {
	var input struct {
		Name                   string `json:"Name"`
		PhoneNumber            string `json:"PhoneNumber"`
		Email                  string `json:"Email"`
		DateOfBirth            string `json:"DateOfBirth"`
		Gender                 string `json:"Gender"`
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

	// Ambil user ID dari session atau token yang disimpan saat login
	userID := c.MustGet("userID").(uint) // Pastikan kamu menyimpan userID di session atau token

	var user models.User
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Ini untuk update name, phone number, dan email jika disediakan
	if input.Name != "" {
		user.Name = input.Name
	}

	if input.Email != "" {
		user.Email = input.Email
	}

	if input.PhoneNumber != "" {
		// Validasi bahwa PhoneNumber hanya berisi angka dan panjangnya minimal 10 digit
		match, _ := regexp.MatchString(`^[0-9]+$`, input.PhoneNumber)
		if !match {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number. Only digits are allowed."})
			return
		}
		user.PhoneNumber = input.PhoneNumber
	}		

	// Parsing DateOfBirth dari string ke time.Time
	if input.DateOfBirth != "" {
		layout := "2006-01-02"
		dob, err := time.Parse(layout, input.DateOfBirth)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
			return
		}
		user.DateOfBirth = models.DateOfBirth(dob)
	}

	// Hanya memperbarui kolom selain password
	user.Gender = input.Gender
	user.EducationalInstitution = input.EducationalInstitution
	user.Profession = input.Profession
	user.Address = input.Address
	user.Province = input.Province
	user.City = input.City
	user.UpdatedAt = time.Now()

	// Simpan perubahan ke database
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return user data yang telah diperbarui
	c.JSON(http.StatusOK, user)
}
