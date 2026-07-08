package test

import (
	"testing"

	"github.com/RafaelFleitas/API-Golang/src/model"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestEncryptPassword(t *testing.T) {
	userDomain := model.NewUserDomain("teste@gmail.com", "senha123@", "teste", int8(22))

	userDomain.EncryptPassword()

	assert.NotEqual(t, "senha123@", userDomain.GetPassword())

	err := bcrypt.CompareHashAndPassword([]byte(userDomain.GetPassword()), []byte("senha123@"))
	assert.NoError(t, err, "O hash deveria corresponder à senha fornecida")

}
