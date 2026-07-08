package test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RafaelFleitas/API-Golang/src/model/repository"
	"github.com/stretchr/testify/assert"
)

func TestFindUserByEmailRepository_Success(t *testing.T) {
	db, mock, err := sqlmock.New() //cria um banco falso
	if err != nil {
		t.Fatalf("erro ao criar mock do banco: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "age", "avatar_url"}).AddRow(1, "Teste", "teste@gmail.com", "hash123", int8(25), "")

	mock.ExpectQuery("SELECT id, name, email, password, age, avatar_url FROM users WHERE email").
		WithArgs("teste@gmail.com").
		WillReturnRows(rows)

	userRepository := repository.NewUserRepository(db)

	result, restErr := userRepository.FindUserByEmailRepository("teste@gmail.com")

	assert.Nil(t, restErr)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.GetID())
	assert.Equal(t, "Teste", result.GetName())
	assert.Equal(t, "teste@gmail.com", result.GetEmail())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindUserByEmailRepository_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, name, email, password, age, avatar_url FROM users WHERE email").
		WithArgs("naoexiste@gmail.com").
		WillReturnError(sql.ErrNoRows)

	userRepository := repository.NewUserRepository(db)

	result, restErr := userRepository.FindUserByEmailRepository("naoexiste@gmail.com")

	assert.Nil(t, result)
	assert.NotNil(t, restErr)
	assert.Equal(t, 404, restErr.Code)
}
