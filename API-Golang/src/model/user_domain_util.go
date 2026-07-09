package model

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword substitui a senha em texto puro pelo hash bcrypt antes de salvar no banco
func (ud *userDomain) EncryptPassword() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	ud.password = string(hashedPassword)
}

func HashFileName(input string) string {
	sum := sha256.Sum256([]byte(input))
	return hex.EncodeToString(sum[:])
}
