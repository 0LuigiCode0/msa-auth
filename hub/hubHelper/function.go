package hubHelper

import (
	"fmt"
	"time"

	"github.com/0LuigiCode0/msa-auth/helper"
	"github.com/0LuigiCode0/msa-auth/store/mongo/model"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitHelper(H HandlerForHelper) Helper { return &help{HandlerForHelper: H} }

// GenerateJwt generates new token
func (h *help) GenerateJwt(id primitive.ObjectID) (string, error) {
	claims := &model.UserClaims{
		ID:   id,
		Time: time.Now(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(helper.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return tokenString, nil
}
