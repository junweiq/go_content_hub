package main

import (
	"fmt"
	"go_content_hub/internal/api"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.CmsRouter(r)

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
