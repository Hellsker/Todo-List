package v1

import "github.com/gin-gonic/gin"

type message struct {
	Message string `json:"message" example:"Test message"`
}

func messageResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, message{msg})
}
