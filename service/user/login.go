package user

import (
	"Cross-field-shop/model"
	"Cross-field-shop/pkg/token"
	"Cross-field-shop/util"
	"errors"
	"github.com/jinzhu/gorm"
)

func Login(account, accountPassword string) (string, int, error) {
	user, err := model.GetUserByEmailAndPassword(account, accountPassword)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", 0, errors.New("user doesn't exist")
		}
		return "", 0, err
	}

	// 生成 auth token
	tokenString, err := token.GenerateToken(&token.TokenPayload{
		ID:      user.Id,
		Expired: util.GetExpiredTime(),
	})
	if err != nil {
		return "", 0, err
	}

	return tokenString, int(user.Id), nil
}
