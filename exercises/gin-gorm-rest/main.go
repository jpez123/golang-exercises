package main

import (
	"github.com/gin-gonic/gin"
	"github.com/restapi/controllers"
	"github.com/restapi/models"
)

func main() {
	//Router
	r := gin.Default()

	//Setup models/schemas
	db := models.SetupModels()

	//Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	//CRUD
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	//Run server
	r.Run()
}
