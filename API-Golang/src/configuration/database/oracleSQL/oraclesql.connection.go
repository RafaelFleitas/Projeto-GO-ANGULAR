package oraclesql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	_ "github.com/sijms/go-ora/v2"
)

var (
	ORACLE_URL = "ORACLE_URL"
)

func NewOracleConnection() (*sql.DB, error) {

	connStr := os.Getenv(ORACLE_URL)
	if connStr == "" {
		return nil, fmt.Errorf("Variável/variáveis de ambiente não configuradas")
	}

	db, err := sql.Open("oracle", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logger.Info("Conseguiu se conectar")

	return db, nil

}
