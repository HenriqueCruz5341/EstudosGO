package repository

import (
	"context"
	"crud-api/src/configuration/logger"
	"crud-api/src/configuration/rest_err"
	"crud-api/src/model"
	"crud-api/src/model/repository/entity/converter"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository",
		zap.String("journey", "createUser"))

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create user",
			err,
			zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.Id = result.InsertedID.(primitive.ObjectID)
	return converter.ConvertEntityToDomain(*value), nil
}
