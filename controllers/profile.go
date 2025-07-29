package controllers

import (
	"linkmatch-be/responses"

	"github.com/gin-gonic/gin"
)

func Profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetString("username")
		responses.JSONSuccess(c.Writer, gin.H{
			"username": username,
			"message":  "You have access to a protected route",
		}, nil, nil)
	}
}
