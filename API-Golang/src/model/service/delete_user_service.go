package service

import (
	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserService(userId int64) *rest_err.RestErr {
	logger.Info("Init deleteUserService model",
		zap.String("journey", "deleteUserService"),
	)

	err := ud.userRepository.DeleteUserRepository(userId)
	if err != nil {
		logger.Error("Error trying to call user", err,
			zap.String("journey", "deleteUserService"),
		)
		return nil
	}

	logger.Info("deleteUserService executed successfully",
		zap.String("journey", "deleteUserService"))

	return nil
}
