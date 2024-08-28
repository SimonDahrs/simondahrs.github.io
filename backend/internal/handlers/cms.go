package handlers

import (
	"backend/internal/services"
	"net/http"
	"github.com/gin-gonic/gin"
)
func GetAllContent(c *gin.Context) {
	posts := services.RetrievePosts()
	c.JSON(http.StatusOK, posts)
}