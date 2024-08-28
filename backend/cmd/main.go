package main

import (
    "github.com/gin-gonic/gin"
	"backend/internal/handlers"
)

func main() {
	r := gin.Default()
	r.GET("/content", handlers.GetAllContent)
	r.Run("127.0.0.1:8080")
}