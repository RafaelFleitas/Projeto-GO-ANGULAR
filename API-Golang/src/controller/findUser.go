package controller

import (
	"net/http"
	"net/mail"
	"strconv"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// O gin.Context tem todas as informações da request
// Controller dos Find
func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init FindUserByIdController",
		zap.String("journey", "FindUserById"),
	)

	userId := c.Param("userId")
	userIdInt, parseErr := strconv.ParseInt(userId, 10, 64)

	if parseErr != nil {
		logger.Error("Error trying to validate user id", parseErr,
			zap.String("journey", "FindUserById"),
		)
		errorMessage := rest_err.NewBadRequestError("Invalid user id")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIdService(userIdInt)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByIdController successfully executed",
		zap.String("journey", "FindUserById"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {

	logger.Info("Init FindUserByEmailController",
		zap.String("journey", "FindUserById"),
	)

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate user email", err,
			zap.String("journey", "FindUserByEmail"),
		)

		errorMessage := rest_err.NewBadRequestError("Invalid user email")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailService(userEmail)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmailController successfully executed",
		zap.String("journey", "FindUserByEmail"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}
