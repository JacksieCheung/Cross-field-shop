package middleware

import (
	"Cross-field-shop/handler"
	"Cross-field-shop/pkg/auth"
	"Cross-field-shop/pkg/errno"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware ... 认证中间件
// AuthMiddleware ... 认证中间件
// limit 为限制的权限等级
func AuthMiddleware(limit uint32) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		ctx, err := auth.ParseRequest(c)
		if err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, err.Error())
			c.Abort()
			return
		} else if ctx.Role < limit { // 原来是与运算，这里简单设置成 大于小于运算。
			handler.SendResponse(c, errno.ErrPermissionDenied, "")
			c.Abort()
			return
		}

		c.Set("userID", ctx.ID)
		c.Set("role", ctx.Role)
		c.Set("expiresAt", ctx.ExpiresAt)

		c.Next()
	}
}
