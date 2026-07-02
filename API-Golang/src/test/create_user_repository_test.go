package test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

const ORACLE_URL = "ORACLE_URL"

func TestUserRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO users").
		WithArgs("Teste", "teste@gmail.com", "hash123", 25).
		WillReturnResult(sqlmock.NewResult(1, 1)) // id=1, 1 linha afetada

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectativas não atendidas: %v", err)
	}
}
