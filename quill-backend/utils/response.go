package utils

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(200, gin.H{
		"message": message,
		"data":    data,
	})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error": message,
	})
}
