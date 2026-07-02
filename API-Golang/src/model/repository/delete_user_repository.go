package repository

import (
	"context"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
)

func (ur *userRepository) DeleteUserRepository(userId int64) *rest_err.RestErr {

	logger.Info("Init deleteUser repository")

	_, err := ur.databaseConnection.ExecContext(
		context.Background(),
		"DELETE FROM users WHERE id = :1",
		userId,
	)

	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil

}
