package middleware

import (
	"Data-acquisition-subsystem/handler"
	"Data-acquisition-subsystem/pkg/auth"
	"Data-acquisition-subsystem/pkg/errno"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware ... 认证中间件
func AuthMiddleware(c *gin.Context) {
	// Parse the json web token.
	ctx, err := auth.ParseRequest(c)
	if err != nil {
		handler.SendResponse(c, errno.ErrTokenInvalid, err.Error())
		c.Abort()
		return
	}

	c.Set("userID", ctx.ID)
	c.Set("expiresAt", ctx.ExpiresAt)

	c.Next()
}
