package service

import (
	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserService(userId int64, userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateUserService model",
		zap.String("journey", "UpdateUserService"),
	)

	err := ud.userRepository.UpdateUserRepository(userId, userDomain)

	if err != nil {
		logger.Error("Error trying to call user", err,
			zap.String("journey", "updateUserService"),
		)
		return nil, err
	}

	updatedUser, err := ud.userRepository.FindUserByIdRepository(userId)

	if err != nil {
		logger.Error("Error trying to find updated user", err,
			zap.String("journey", "updateUserService"),
		)
		return nil, err
	}

	logger.Info("UpdateUserService service executed successfully",
		zap.String("journey", "updateUserService"))

	return updatedUser, nil
}
