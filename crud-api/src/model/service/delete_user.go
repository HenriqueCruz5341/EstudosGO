package service

import (
	"crud-api/src/configuration/logger"
	"crud-api/src/configuration/rest_err"

	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Init deleteUser model", zap.String("journey", "deleteUser"))
	return nil
}
