package controllers

import (
	"log"
	"net/http"
	"time"
	"ukom_crud/config"
	"ukom_crud/models"
	"ukom_crud/utils" // Import utils
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Struktur untuk input registrasi
type RegistrasiInput struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

// Registrasi User
func RegistrasiUser(c *gin.Context) {
	var input RegistrasiInput

	// Bind JSON input ke struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi PhoneNumber hanya angka
	if !utils.IsNumeric(input.PhoneNumber) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number must contain only digits"})
		return
	}

	// Hashing password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Membuat user baru
	user := models.User{
		Name:        input.Name,
		Email:       input.Email,
		Password:    string(hashedPassword),
		PhoneNumber: input.PhoneNumber,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Simpan user ke database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()})
		return
	}

	// Registrasi sukses
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Struktur untuk input login
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Login User
func LoginUser(c *gin.Context) {
	var input LoginInput

	// Bind JSON input ke struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cari user berdasarkan email
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		log.Println("Error finding user:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password is incorrect"})
		return
	}

	// Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password is incorrect"})
		return
	}

	// Simpan user ID ke session
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	// Login sukses
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// Logout User
func LogoutUser(c *gin.Context) {
	// Hapus session
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	// Logout sukses
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

// // Middleware untuk memeriksa apakah user sudah login
// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Cek session
// 		session := sessions.Default(c)
// 		userID := session.Get("user_id")

// 		// Jika user belum login
// 		if userID == nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not logged in"})
// 			c.Abort()
// 			return
// 		}

// 		// Lanjut ke handler berikutnya jika user sudah login
// 		c.Next()
// 	}
// }
