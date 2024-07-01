package repository

import (
	"crud-api/src/configuration/rest_err"
	"crud-api/src/model"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		databaseConnection: database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}
