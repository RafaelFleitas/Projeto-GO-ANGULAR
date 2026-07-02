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

func (uc *userControllerInterface) LoginUser(c *gin.Context) {

	logger.Info("Init loginUser controller",
		zap.String("journey", "loginUser"),
	)

	var userLogin request.UserLogin // Declara uma variável do tipo UserLogin para armazenar os dados da requisição

	// ShouldBindJson garante que os dados recebidos estão no formato correto
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		logger.Error("Init loginUser controller", err,
			zap.String("journey", "loginUser"),
		)
		restErr := validation.ValidateUserError(err) //Chama a função de validação de erros em validation

		c.JSON(restErr.Code, restErr) // Retorna uma resposta JSON com o status HTTP do erro e o objeto de erro.
		return
	}

	domain := model.NewUserLoginDomain(
		userLogin.Email,
		userLogin.Password,
	)

	domainResult, token, err := uc.service.LoginUserService(domain)
	if err != nil {
		logger.Error("Error trying to call loginUser service", err,
			zap.String("journey", "loginUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User login successfully",
		zap.String("journey", "loginUser"),
	)
	c.Header("Authorization", token)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
