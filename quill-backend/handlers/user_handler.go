package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Note created successfully",
	})
}

func CreateNote(c *gin.Context) {
	var note models.Note
}