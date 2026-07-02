package converter

import (
	"github.com/RafaelFleitas/API-Golang/src/model"
	"github.com/RafaelFleitas/API-Golang/src/model/repository/entity"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}

func ConvertEntityToDomain(userEntity *entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		userEntity.Email,
		userEntity.Password,
		userEntity.Name,
		userEntity.Age,
	)

	domain.SetID(userEntity.ID)

	return domain

}
