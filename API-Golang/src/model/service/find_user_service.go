package service

import (
	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIdService(id int64) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByIdService services",
		zap.String("journey", "FindUserById"),
	)

	return ud.userRepository.FindUserByIdRepository(id)
}

func (ud *userDomainService) FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmailService services",
		zap.String("journey", "FindUserById"),
	)

	return ud.userRepository.FindUserByEmailRepository(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordService(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmailAndPasswordService services",
		zap.String("journey", "FindUserById"),
	)

	return ud.userRepository.FindUserByEmailAndPasswordRepository(email, password)
}
