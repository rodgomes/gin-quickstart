package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: check authentication and fail and credentials missing
		c.Next()
	}
}
