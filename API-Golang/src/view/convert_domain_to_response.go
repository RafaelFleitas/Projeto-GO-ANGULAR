package view

import (
	"github.com/RafaelFleitas/API-Golang/src/controller/model/response"
	"github.com/RafaelFleitas/API-Golang/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
