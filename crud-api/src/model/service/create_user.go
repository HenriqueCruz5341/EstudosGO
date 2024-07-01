package service

import (
	"crud-api/src/configuration/logger"
	"crud-api/src/configuration/rest_err"
	"crud-api/src/model"

	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}
	return userDomainRepository, nil
}
