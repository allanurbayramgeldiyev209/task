package helpers

import "github.com/gin-gonic/gin"

func HandleError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"status":  false,
		"message": message,
	})
}
