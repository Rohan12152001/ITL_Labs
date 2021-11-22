package middleware

import (
	"github.com/gin-gonic/gin"
)

type AuthManager interface {
	AuthMiddleWareWithUser(c *gin.Context)
}