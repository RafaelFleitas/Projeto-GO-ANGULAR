package service

import (
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"github.com/RafaelFleitas/API-Golang/src/model/repository"
)

// userDomainService é a camada de serviço. Ela fica entre o controller e o repositório,
// e é onde ficam as regras de negócio (ex: criptografar senha antes de salvar).
type userDomainService struct {
	userRepository repository.UserRepository
}

// NewUserDomainService recebe o repositório e devolve o serviço pronto para uso
func NewUserDomainService(repository repository.UserRepository) UserDomainService {
	return &userDomainService{repository}
}

// UserDomainService define quais operações de usuário existem na aplicação.
// O controller só enxerga essa interface, nunca o repositório diretamente.
type UserDomainService interface {
	CreateUserService(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserService(userId int64, userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByIdService(id int64) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr)
	LoginUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr)
	DeleteUserService(int64) *rest_err.RestErr
}
