package v1

import "github.com/gin-gonic/gin"

type error struct {
	Error string `json:"error" example:"Test Error!"`
}

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, error{msg})
}
