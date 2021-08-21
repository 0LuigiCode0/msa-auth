package model

import (
	"time"

	"github.com/0LuigiCode0/msa-auth/helper"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserClaims struct {
	jwt.StandardClaims
	ID   primitive.ObjectID
	Time time.Time
}

// Модель пользователя в БД
type UserModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Login    string             `bson:"login" json:"login"`
	Password string             `bson:"password" json:"password"`
	Role     helper.Role        `bson:"role" json:"role"`
}

type RequsetAuthGRPC struct {
	Jwt   string        `json:"jwt"`
	Roles []helper.Role `json:"role"`
}
