package repository

import (
	"context"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model"
)

func (ur *userRepository) UpdateUserRepository(
	userId int64, userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init updateUser repository")

	_, err := ur.databaseConnection.ExecContext(
		context.Background(),
		"UPDATE users SET name = :1, age = :2 WHERE id = :3",
		userDomain.GetName(),
		userDomain.GetAge(),
		userId,
	)

	if err != nil {
		logger.Error("Error trying to update user", err)
		return rest_err.NewInternalServerError("Erro interno, tente novmaente")
	}

	return nil

}
