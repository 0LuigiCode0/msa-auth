package store

import (
	"x-msa-auth/helper"
	"x-msa-auth/store/mongo/model"
	corehelper "x-msa-core/helper"

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
	err := s.db.Collection(string(helper.CollUsers)).FindOne(corehelper.Ctx, primitive.M{
		"_id": id,
	}).Decode(user)
	return user, err
}

func (s *userStore) SelectByLogin(login string) (*model.UserModel, error) {
	user := &model.UserModel{}
	err := s.db.Collection(string(helper.CollUsers)).FindOne(corehelper.Ctx, primitive.M{
		"login": login,
	}).Decode(user)
	return user, err
}
