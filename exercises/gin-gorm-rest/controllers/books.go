package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/restapi/models"
)

// FindBooks - GET /books
// Get all books
func FindBooks(c *gin.Context) {
	//Get database instance
	db := c.MustGet("db").(*gorm.DB)

	//Get list of books from schema
	var books []models.Book
	db.Find(&books)

	//Returns data as JSON
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// CreateBook - POST /books
// Create new book
func CreateBook(c *gin.Context) {
	//Get database instance
	db := c.MustGet("db").(*gorm.DB)

	//Validate input
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Creates book
	book := models.Book{Title: input.Title, Author: input.Author}
	db.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// FindBook - GET /books/:id
// Finda  book
func FindBook(c *gin.Context) {
	//Get database instance
	db := c.MustGet("db").(*gorm.DB)

	//Get model if exist
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook - PATCH /books/:id
// Updat entry
func UpdateBook(c *gin.Context) {
	//Get database instance
	db := c.MustGet("db").(*gorm.DB)

	//Get model if exist
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	//Validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Update current entry with new input
	db.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook - DELETE /books/:id
// Delete entry
func DeleteBook(c *gin.Context) {
	//Get database instance
	db := c.MustGet("db").(*gorm.DB)

	//Get model if exist
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
