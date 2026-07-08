package test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"github.com/RafaelFleitas/API-Golang/src/model/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO users").
		WithArgs("Teste", "teste@gmail.com", "hash123", int8(25), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(0, 1)) // id=0, 1 linha afetada

	userRepository := repository.NewUserRepository(db)
	userDomain := model.NewUserDomain("Teste", "teste@gmail.com", "hash123", int8(25))

	result, restErr := userRepository.CreateUserRepository(userDomain)

	require.Nil(t, restErr)
	require.NotNil(t, result)
	assert.Equal(t, "Teste", result.GetName())
	assert.NoError(t, mock.ExpectationsWereMet())
}
