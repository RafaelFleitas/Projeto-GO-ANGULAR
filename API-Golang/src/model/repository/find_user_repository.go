package repository

import (
	"context"
	"database/sql"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"github.com/RafaelFleitas/API-Golang/src/model/repository/entity"
	"github.com/RafaelFleitas/API-Golang/src/model/repository/entity/converter"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (ur *userRepository) FindUserByEmailRepository(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindByEmail user repository")

	row := ur.databaseConnection.QueryRowContext(
		context.Background(),
		"SELECT id, name, email, password, age FROM users WHERE email = :1",
		email,
	)

	userEntity := &entity.UserEntity{}

	err := row.Scan(&userEntity.ID, &userEntity.Name, &userEntity.Email, &userEntity.Password, &userEntity.Age)

	if err != nil {
		logger.Error("Error trying to find user by email", err)
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError("user not found")
		}
		return nil, rest_err.NewInternalServerError("Error trying to find user by email")
	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("journey", "FindUserByEmail"),
		zap.String("email: ", email))

	return converter.ConvertEntityToDomain(userEntity), nil

}

func (ur *userRepository) FindUserByIdRepository(id int64) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserById user repository")

	row := ur.databaseConnection.QueryRowContext(
		context.Background(),
		"SELECT id, name, email, password, age FROM users WHERE ID = :1",
		id,
	)

	userEntity := &entity.UserEntity{}

	err := row.Scan(&userEntity.ID, &userEntity.Name, &userEntity.Email, &userEntity.Password, &userEntity.Age)

	if err != nil {
		logger.Error("Error trying to find user by ID", err)
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError("user not found")
		}
		return nil, rest_err.NewInternalServerError("Error trying to find user by ID")
	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("journey", "FindUserById"),
		zap.Int64("ID: ", id))

	return converter.ConvertEntityToDomain(userEntity), nil

}

func (ur *userRepository) FindUserByEmailAndPasswordRepository(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindByEmailAndPassword user repository")

	row := ur.databaseConnection.QueryRowContext(
		context.Background(),
		"SELECT id, name, email, password, age FROM users WHERE email = :1",
		email,
	)

	userEntity := &entity.UserEntity{}

	err := row.Scan(&userEntity.ID, &userEntity.Name, &userEntity.Email, &userEntity.Password, &userEntity.Age)

	if err != nil {
		logger.Error("Error trying to find user by email and password", err)
		if err == sql.ErrNoRows {
			return nil, rest_err.NewForbiddenError("User or password is invalid")
		}
		return nil, rest_err.NewInternalServerError("Error trying to find user by email and password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(password)); err != nil {
		return nil, rest_err.NewForbiddenError("User or password is invalid")
	}

	logger.Info("FindByEmailAndPassword repository executed successfully",
		zap.String("journey", "FindByEmailAndPassword"),
		zap.String("email: ", email))

	return converter.ConvertEntityToDomain(userEntity), nil

}
