package view

import (
	"fmt"
	"os"

	"github.com/RafaelFleitas/API-Golang/src/controller/model/response"
	"github.com/RafaelFleitas/API-Golang/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:           userDomain.GetID(),
		Email:        userDomain.GetEmail(),
		Name:         userDomain.GetName(),
		Age:          userDomain.GetAge(),
		ProfileImage: buildAvatarURL(userDomain.GetAvatarURL()),
	}
}

func buildAvatarURL(avatarPath string) string {
	if avatarPath == "" {
		return ""
	}
	baseURL := os.Getenv("API_BASE_URL")
	return fmt.Sprintf("%s%s", baseURL, avatarPath)
}

func ConvertDomainListToResponse(usersDomain []model.UserDomainInterface) []response.UserResponse {
	usersResponse := []response.UserResponse{}

	for _, userDomain := range usersDomain {
		usersResponse = append(usersResponse, ConvertDomainToResponse(userDomain))
	}

	return usersResponse
}
