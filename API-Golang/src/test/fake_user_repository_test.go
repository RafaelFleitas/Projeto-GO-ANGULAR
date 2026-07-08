package test

import (
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"github.com/RafaelFleitas/API-Golang/src/model/repository"
)

type FakeUserRepository struct {
	repository.UserRepository // embutido como nil — dá o "de resto" da interface de graça

	CreateUserFunc      func(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailFunc func(string) (model.UserDomainInterface, *rest_err.RestErr)
}

func (f *FakeUserRepository) CreateUserRepository(u model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	return f.CreateUserFunc(u)
}

func (f *FakeUserRepository) FindUserByEmailRepository(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	return f.FindUserByEmailFunc(email)
}
