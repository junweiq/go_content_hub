package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "9941"
	}

	err := r.Run(":" + port)
	if err != nil {
		fmt.Printf("r run error = %v", err)
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
