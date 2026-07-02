package repository

import (
	"database/sql"

	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model"
)

// userRepository guarda a conexão com o banco. Todos os métodos de acesso ao banco ficam nessa struct.
type userRepository struct {
	databaseConnection *sql.DB
}

// UserRepository define quais operações de banco estão disponíveis para o usuário.
// Novos métodos (FindUser, UpdateUser, DeleteUser) serão adicionados aqui conforme forem implementados.
type UserRepository interface {
	CreateUserRepository(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByIdRepository(id int64) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailRepository(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailAndPasswordRepository(email, password string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserRepository(int64, model.UserDomainInterface) *rest_err.RestErr
	DeleteUserRepository(int64) *rest_err.RestErr
}

// NewUserRepository recebe a conexão com o banco e devolve um repositório pronto para uso
func NewUserRepository(database *sql.DB) UserRepository {
	return &userRepository{
		database,
	}
}
