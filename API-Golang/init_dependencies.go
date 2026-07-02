package main

import (
	"database/sql"

	"github.com/RafaelFleitas/API-Golang/src/controller"
	"github.com/RafaelFleitas/API-Golang/src/model/repository"
	"github.com/RafaelFleitas/API-Golang/src/model/service"
)

func initDependencies(db *sql.DB) controller.UserControllerInterface {

	// Inicializa o router, registra as rotas da aplicação e inicia o servidor na porta 8000
	userRepository := repository.NewUserRepository(db)
	service := service.NewUserDomainService(userRepository)

	return controller.NewUserControllerInterface(service)

}
