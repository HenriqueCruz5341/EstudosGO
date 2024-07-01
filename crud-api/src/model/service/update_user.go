package service

import (
	"crud-api/src/configuration/logger"
	"crud-api/src/configuration/rest_err"
	"crud-api/src/model"

	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))
	return nil
}
