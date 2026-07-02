package service

import (
	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {

	logger.Info("Init LoginUser model",
		zap.String("journey", "LoginUser"),
	)

	user, err := ud.findUserByEmailAndPasswordService(userDomain.GetEmail(), userDomain.GetPassword())

	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()

	if err != nil {
		return nil, "", err
	}

	logger.Info("LoginUser service executed successfully",
		zap.String("journey", "LoginUser"))

	return user, token, nil
}
