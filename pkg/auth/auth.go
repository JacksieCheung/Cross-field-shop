package auth

import (
	"Cross-field-shop/pkg/errno"
	"Cross-field-shop/pkg/token"
	"github.com/gin-gonic/gin"
)

// Context is the context of the JSON web token.
type Context struct {
	ID        uint32
	Role      uint32
	ExpiresAt int64 // 过期时间（时间戳，10位）
}

// Parse parses the token, and returns the context if the token is valid.
func Parse(tokenString string) (*Context, error) {
	t, err := token.ResolveToken(tokenString)
	if err != nil {
		return nil, err
	}

	return &Context{
		ID:        t.ID,
		Role:      t.Role,
		ExpiresAt: t.ExpiresAt,
	}, nil
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")
	if len(header) == 0 {
		return nil, errno.ErrMissingHeader
	}

	return Parse(header)
}
