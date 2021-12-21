package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rahmanfadhil/gin-bookstore/controllers"
	"github.com/rahmanfadhil/gin-bookstore/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run(":4000")
}