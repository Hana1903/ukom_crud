package controllers

import (
	"net/http"
	"time"
	"ukom_crud/config"
	"ukom_crud/models"
	"github.com/gin-gonic/gin"
)

// Create Order
func CreateOrder(c *gin.Context) {
	var input struct {
		UserID       int64   `json:"user_id"`
		PackageID    int64   `json:"package_id"`
		PaymentStatus string  `json:"payment_status"`
		TotalPrice   float64 `json:"total_price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create new Order
	order := models.Order{
		UserID:       input.UserID,
		PackageID:    input.PackageID,
		PaymentStatus: input.PaymentStatus,
		TotalPrice:   input.TotalPrice,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Save order to the database
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// Get all orders
func GetOrders(c *gin.Context) {
	var orders []models.Order
	if err := config.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// Get a single order by ID
func GetOrderByID(c *gin.Context) {
	var order models.Order
	if err := config.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// Update order by ID
func UpdateOrder(c *gin.Context) {
	var order models.Order
	if err := config.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var input struct {
		UserID       int64   `json:"user_id"`
		PackageID    int64   `json:"package_id"`
		PaymentStatus string  `json:"payment_status"`
		TotalPrice   float64 `json:"total_price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update order details
	order.UserID = input.UserID
	order.PackageID = input.PackageID
	order.PaymentStatus = input.PaymentStatus
	order.TotalPrice = input.TotalPrice
	order.UpdatedAt = time.Now()

	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// Delete order by ID
func DeleteOrder(c *gin.Context) {
	var order models.Order
	if err := config.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := config.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
