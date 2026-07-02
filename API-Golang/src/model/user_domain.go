package model

// userDomain é a struct privada que guarda os dados do usuário.
// Os campos são privados para que só sejam acessados pelos métodos abaixo.
type userDomain struct {
	id       int64
	email    string
	password string
	name     string
	age      int8
}

// Getters/Setters — única forma de ler os dados do usuário fora desse pacote
func (ud *userDomain) GetEmail() string    { return ud.email }
func (ud *userDomain) GetPassword() string { return ud.password }
func (ud *userDomain) GetName() string     { return ud.name }
func (ud *userDomain) GetAge() int8        { return ud.age }
func (ud *userDomain) GetID() int64        { return ud.id }
func (ud *userDomain) SetID(id int64)      { ud.id = id }
