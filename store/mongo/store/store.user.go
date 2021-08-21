package store

import (
	coreHelper "github.com/0LuigiCode0/msa-core/helper"

	"github.com/0LuigiCode0/msa-auth/helper"
	"github.com/0LuigiCode0/msa-auth/store/mongo/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Store хранилище
type UserStore interface {
	SelectByID(id primitive.ObjectID) (*model.UserModel, error)
	SelectByLogin(login string) (*model.UserModel, error)
}

//Store хранилище
type userStore struct {
	db *mongo.Database
}

func InitUserStore(db *mongo.Database) UserStore {
	return &userStore{db: db}
}

func (s *userStore) SelectByID(id primitive.ObjectID) (*model.UserModel, error) {
	user := &model.UserModel{}
	err := s.db.Collection(string(helper.CollUsers)).FindOne(coreHelper.Ctx, primitive.M{
		"_id": id,
	}).Decode(user)
	return user, err
}

func (s *userStore) SelectByLogin(login string) (*model.UserModel, error) {
	user := &model.UserModel{}
	err := s.db.Collection(string(helper.CollUsers)).FindOne(coreHelper.Ctx, primitive.M{
		"login": login,
	}).Decode(user)
	return user, err
}
