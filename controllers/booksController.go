package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func BooksCreate(c *gin.Context) {
	// post req body
	var body struct{
		Title  string
		Author string
		Pages  int
	}

	c.Bind(&body)
	
	//create
	book := models.Book{Title: body.Title, Author: body.Author, Pages:body.Pages}

	result := initializers.DB.Create(&book);

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return
	c.JSON(200, gin.H{
		"book": book,
	})
}

func BooksRead(c *gin.Context) {
	// create array
	var books []models.Book

	// use gorm and find objects from DB
	initializers.DB.Find(&books)

	//respos 
	c.JSON(200, gin.H{
		"books": books,
	})
}

//  GET :id <-
func BooksReadOne(c *gin.Context) {
	//get id from dynamic get req
	id := c.Param("id")

	// create variable
	var book models.Book

	initializers.DB.First(&book,id)

	c.JSON(200, gin.H{
		"book": book,
	})


}

// POST :id ->
func BooksUpdate(c *gin.Context) {

	id := c.Param("id")

	// post req body
	var body struct{
		Title  string
		Author string
		Pages  int
	}

	c.Bind(&body)

	var book models.Book
	initializers.DB.First(&book, id)

	initializers.DB.Model(&book).Updates(models.Book{
		Title: body.Title,
		Author: body.Author,
		Pages: body.Pages,
	})

	c.JSON(200, gin.H{
		"book": book,
	})
	
}

func BooksDelete(c *gin.Context){

	id := c.Param("id")

	initializers.DB.Delete(&models.Book{}, id)

	c.Status(200)


}