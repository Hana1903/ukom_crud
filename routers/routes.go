package routers

import (
	"ukom_crud/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) *gin.Engine {
	// router := gin.Default()

	router.POST("/register", controllers.RegistrasiUser)
	router.POST("/login", controllers.LoginUser)
	router.POST("/logout", controllers.LogoutUser)

	router.GET("/packages/:id/questions", controllers.GetPackageQuestions)

	// User Routes
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUserByID)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	// Routes Profile
	router.PUT("//users/profile", controllers.UpdateUserProfile)

	router.POST("/exams", controllers.CreateExam)        
	router.GET("/exams", controllers.GetExams)           
	router.GET("/exams/:id", controllers.GetExamByID)     
	router.PUT("/exams/:id", controllers.UpdateExam)      
	router.DELETE("/exams:id", controllers.DeleteExam)
	
	// Question routes
	router.GET("/questions", controllers.GetQuestions)
	router.GET("/questions/:id", controllers.GetQuestionByID)
	router.POST("/questions", controllers.CreateQuestion)
	router.PUT("/questions/:id", controllers.UpdateQuestion)
	router.DELETE("/questions/:id", controllers.DeleteQuestion)

	// Packet routes
	router.GET("/packages", controllers.GetPackages)
	router.GET("/packages/:id", controllers.GetPackageByID)
	router.POST("/packages", controllers.CreatePackage)
	router.PUT("/packages/:id", controllers.UpdatePackage)
	router.DELETE("/packages/:id", controllers.DeletePackage)

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrders)
	router.GET("/orders/:id", controllers.GetOrderByID)
	router.PUT("/orders/:id", controllers.UpdateOrder)
	router.DELETE("/orders/:id", controllers.DeleteOrder)

	router.POST("/exam_questions", controllers.CreateExamQuestion)
	router.GET("/exam_questions", controllers.GetExamQuestions)
	router.GET("/exam_questions/:id", controllers.GetExamQuestionByID)
	router.PUT("/exam_questions/:id", controllers.UpdateExamQuestion)
	router.DELETE("/exam_questions/:id", controllers.DeleteExamQuestion)
	
	return router
}