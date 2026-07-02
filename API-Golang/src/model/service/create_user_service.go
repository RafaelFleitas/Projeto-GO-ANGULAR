package service

import (
	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init CreateUser model",
		zap.String("journey", "createUser"),
	)

	//Verifica se o email já está sendo utilizado por outra conta
	user, _ := ud.FindUserByEmailService(userDomain.GetEmail())
	if user != nil {
		return nil, rest_err.NewBadRequestError("User email already exists")
	}

	// Criptografa a senha antes de passar para o repositório salvar no banco
	userDomain.EncryptPassword()

	// Repassa para o repositório que vai executar o INSERT no Oracle
	userDomainRepository, err := ud.userRepository.CreateUserRepository(userDomain)

	if err != nil {
		logger.Error("Error trying to call user", err,
			zap.String("journey", "createUser"),
		)
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("journey", "createUser"))

	return userDomainRepository, nil
}
