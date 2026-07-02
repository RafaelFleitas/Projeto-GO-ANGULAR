package model

import "github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"

// UserDomainInterface define o que um usuário precisa ter na aplicação.
// Toda a aplicação usa essa interface — nunca a struct diretamente.
// Isso protege os campos privados e facilita trocar a implementação no futuro.
type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetID() int64
	SetID(id int64)

	EncryptPassword()
	GenerateToken() (string, *rest_err.RestErr)
}

// NewUserDomain é o construtor do usuário. Recebe os dados da requisição e devolve a interface.
func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserUpdateDomain(name string, age int8) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}

func NewUserLoginDomain(email, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}
