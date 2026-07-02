package controller

import (
	"net/http"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/validation"
	"github.com/RafaelFleitas/API-Golang/src/controller/model/request"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"github.com/RafaelFleitas/API-Golang/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Controller do Create
func (uc *userControllerInterface) CreateUser(c *gin.Context) {

	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest // Declara uma variável do tipo UserRequest para armazenar os dados da requisição

	// ShouldBindJson garante que os dados recebidos estão no formato correto
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Init CreateUser controller", err,
			zap.String("journey", "createUser"),
		)
		restErr := validation.ValidateUserError(err) //Chama a função de validação de erros em validation

		c.JSON(restErr.Code, restErr) // Retorna uma resposta JSON com o status HTTP do erro e o objeto de erro.
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.service.CreateUserService(domain)
	if err != nil {
		logger.Error("Error trying to call CreateUser service", err,
			zap.String("journey", "createUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
