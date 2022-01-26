package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}
