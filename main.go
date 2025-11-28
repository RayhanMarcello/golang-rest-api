package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	// get methode
	router.GET("/", rootHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)

	// post methode
	router.POST("/books", postBooksHandler)

	if err := router.Run(":8000");
	err != nil {
		log.Fatalf("failed to rn server: %v", err)
	}
}

// handler function
func rootHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
			"message" : "testing endpoint root dengan methode GET",
		})
}

func booksHandler(c *gin.Context){
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id" : id,
		"title" : title,
	})
}

func queryHandler(c *gin.Context){
	query := c.Query("query")
	c.JSON(http.StatusOK, gin.H{
		"query" : query,
	})
}

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price int `json:"price" binding:"required,number"`
}

func postBooksHandler(c *gin.Context){
	var bookInput BookInput
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		log.Fatalln("error", err)
	}else{
		c.JSON(http.StatusOK, gin.H{
			"Title" : bookInput.Title,
			"Price" : bookInput.Price,			
		})
	}
}