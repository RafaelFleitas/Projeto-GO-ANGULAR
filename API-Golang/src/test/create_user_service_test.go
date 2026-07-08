package test

import (
	"testing"

	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"github.com/RafaelFleitas/API-Golang/src/model/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUserService_Success(t *testing.T) {
	fakeRepo := &FakeUserRepository{
		FindUserByEmailFunc: func(email string) (model.UserDomainInterface, *rest_err.RestErr) {
			return nil, nil // ninguém usa esse email ainda
		},
		CreateUserFunc: func(u model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
			u.SetID(1) // simula o banco gerando o ID
			return u, nil
		},
	}

	userService := service.NewUserDomainService(fakeRepo)
	userDomain := model.NewUserDomain("novo@gmail.com", "senha123", "Novo Usuario", 30)

	result, restErr := userService.CreateUserService(userDomain)

	require.Nil(t, restErr)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.GetID())
	assert.NotEqual(t, "senha123", result.GetPassword()) // senha tem que estar criptografada, não em texto puro
}

func TestCreateUserService_EmailAlreadyExists(t *testing.T) {
	createWasCalled := false

	fakeRepo := &FakeUserRepository{
		FindUserByEmailFunc: func(email string) (model.UserDomainInterface, *rest_err.RestErr) {
			existingUser := model.NewUserDomain(email, "outrasenha", "Outro Usuario", 25)
			return existingUser, nil
		},
		CreateUserFunc: func(u model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
			createWasCalled = true
			return u, nil
		},
	}

	userService := service.NewUserDomainService(fakeRepo)
	userDomain := model.NewUserDomain("jaexiste@gmail.com", "senha123", "Novo Usuario", 30)

	result, restErr := userService.CreateUserService(userDomain)

	require.NotNil(t, restErr)
	assert.Nil(t, result)
	assert.Equal(t, 400, restErr.Code)
	assert.False(t, createWasCalled, "CreateUserRepository não deveria ter sido chamado")

}

func TestCreateUserService_RepositoryError(t *testing.T) {
	fakeRepo := &FakeUserRepository{
		FindUserByEmailFunc: func(email string) (model.UserDomainInterface, *rest_err.RestErr) {
			return nil, nil
		},
		CreateUserFunc: func(u model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
			return nil, rest_err.NewInternalServerError("Erro interno, tente novamente")
		},
	}

	userService := service.NewUserDomainService(fakeRepo)
	userDomain := model.NewUserDomain("novo@gmail.com", "senha123", "Novo Usuario", 30)

	result, restErr := userService.CreateUserService(userDomain)

	assert.Nil(t, result)
	require.NotNil(t, restErr)
	assert.Equal(t, 500, restErr.Code)
}
