package controllers

import (
	"net/http"
	"ukom_crud/config"
	"ukom_crud/models"
	"github.com/gin-gonic/gin"
)

// Get all packages
func GetPackages(c *gin.Context) {
	var packages []models.Package
	if err := config.DB.Find(&packages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve packages."})
		return
	}
	c.JSON(http.StatusOK, packages)
}

// Get package by ID
func GetPackageByID(c *gin.Context) {
	var pkg models.Package
	id := c.Param("id")

	if err := config.DB.First(&pkg, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package not found!"})
		return
	}
	c.JSON(http.StatusOK, pkg)
}

func GetPackageQuestions(c *gin.Context)  {
	var pkg models.Package
	packageID := c.Param("id")

	//Cari package berdasarkan id dan memuat pertanyaan 
	if err := config.DB.Preload("Questions").First(&pkg, packageID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package not found"})
		return
	}
	//Kirim respons dengan pertanyaan dan package
	c.JSON(http.StatusOK, pkg)
}

// Create new package
func CreatePackage(c *gin.Context) {
	var input models.Package
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data!"})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create package."})
		return
	}
	c.JSON(http.StatusOK, input)
}

// Update package by ID
func UpdatePackage(c *gin.Context) {
	var pkg models.Package
	id := c.Param("id")

	if err := config.DB.First(&pkg, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package not found!"})
		return
	}

	var input models.Package
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data!"})
		return
	}

	// Ensure only non-zero fields are updated
	if err := config.DB.Model(&pkg).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update package."})
		return
	}
	c.JSON(http.StatusOK, pkg)
}

// Delete package by ID
func DeletePackage(c *gin.Context) {
	var pkg models.Package
	id := c.Param("id")

	if err := config.DB.First(&pkg, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package not found!"})
		return
	}

	if err := config.DB.Delete(&pkg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete package."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Package deleted successfully!"})
}
