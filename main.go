package main

import (
	"ukom_crud/config"
	"ukom_crud/models"
	"ukom_crud/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect ke database
	config.ConnectDB()

	// Auto Migrate untuk memastikan table terbentuk sesuai models
	config.DB.AutoMigrate(&models.User{}, &models.Exam{}, &models.Package{}, &models.Order{}, &models.Question{}, &models.ExamQuestion{})

	// Inisialisasi router Gin
	router := gin.Default()

	// Pengaturan penyimpanan session di cookie
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// Setup routes
	router = routers.SetupRoutes(router)

	// Menjalankan server di port 8080
	router.Run(":8080")
}

