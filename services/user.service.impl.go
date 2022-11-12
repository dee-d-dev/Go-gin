package services

import (
	"context"

	"github.com/dee-d-dev/go-gin/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx Context) UserService{
	return &UserServiceImpl{
		usercollection: usercollection
		ctx: ctx
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)

	return err
} 

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{key:"user_name", value: name}}
	err := uc.usercollection.FindOne(u.ctx, query).Decode(&user)

	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	filter := bson.D{bson.E{Key:"user_name", Value: user.Name}}

	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key:"user_name", bson.E{}}}}}
	return nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	return nil
}
