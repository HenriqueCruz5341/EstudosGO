package service

import (
	"crud-api/src/configuration/logger"
	"crud-api/src/configuration/rest_err"
	"crud-api/src/model"

	"go.uber.org/zap"
)

func (ud *userDomainService) FindUser(userId string) (*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUser model", zap.String("journey", "findUser"))
	return nil, nil
}
