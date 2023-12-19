package main

import (
	"go-crud/controllers"
	"go-crud/initializers"

	"github.com/gin-gonic/gin"
)


func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	// creating router
	r := gin.Default()

	// creating route /books POST/GET
	r.POST("/books", controllers.BooksCreate)
	r.PUT("/books/:id", controllers.BooksUpdate)
	r.GET("/books", controllers.BooksRead)
	r.GET("/books/:id", controllers.BooksReadOne)
	r.DELETE("/books/:id", controllers.BooksDelete)
	//Listen and run...
	r.Run()
}