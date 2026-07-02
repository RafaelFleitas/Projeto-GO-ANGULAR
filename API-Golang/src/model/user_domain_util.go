package model

import "golang.org/x/crypto/bcrypt"

// EncryptPassword substitui a senha em texto puro pelo hash bcrypt antes de salvar no banco
func (ud *userDomain) EncryptPassword() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	ud.password = string(hashedPassword)
}
