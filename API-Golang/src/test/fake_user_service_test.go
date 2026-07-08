package test

import (
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"github.com/RafaelFleitas/API-Golang/src/model/service"
)

type FakeUserDomainService struct {
	service.UserDomainService
	CreateUserServiceFunc func(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}

func (f *FakeUserDomainService) CreateUserService(u model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	return f.CreateUserServiceFunc(u)
}
