package userdatabase

import (
	"context"

	dbs "github.com/pienaahj/UserDatabase/databases"
	"github.com/pienaahj/UserDatabase/models"
)

const collectionName = "users"

//go:generate mockgen --destination=./mocks/UserDatabase.go github.com/pienaahj/UserDatabase UserDatabase
type UserDatabase interface {
	FindOne(context.Context, interface{}) (*models.User, error)
	Create(context.Context, *models.User) error
	DeleteByUsername(context.Context, string) error
}

type userDatabase struct {
	db dbs.DatabaseHelper
}

func NewUserDatabase(db dbs.DatabaseHelper) UserDatabase {
	return &userDatabase{
		db: db,
	}
}

func (u *userDatabase) FindOne(ctx context.Context, filter interface{}) (*models.User, error) {
	user := &models.User{}
	err := u.db.Collection(collectionName).FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userDatabase) Create(ctx context.Context, usr *models.User) error {
	_, err := u.db.Collection(collectionName).InsertOne(ctx, usr)
	return err
}

func (u *userDatabase) DeleteByUsername(ctx context.Context, username string) error {
	// In this case it is possible to use bson.M{"username":username} but I tend
	// to avoid another dependency in this layer and for demonstration purposes
	// used omitempty in the model
	user := &models.User{
		Username: username,
	}
	_, err := u.db.Collection(collectionName).DeleteOne(ctx, user)
	return err
}
