package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultController(context *gin.Context) {
	context.JSON(http.StatusForbidden, gin.H{
		"error":   "invalid route",
		"message": "list endpoints at /list/api",
	})
}
