package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message" : "testing endpoint root dengan methode GET",
		})
	})

	if err := router.Run();
	err != nil {
		log.Fatalf("failed to rn server: %v", err)
	}
}