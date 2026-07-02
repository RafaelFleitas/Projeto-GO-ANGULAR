package controller

import (
	"net/http"
	"strconv"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/configuration/validation"
	"github.com/RafaelFleitas/API-Golang/src/controller/model/request"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"github.com/RafaelFleitas/API-Golang/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {

	logger.Info("Init UpdateUser controller",
		zap.String("journey", "UpdateUser"),
	)

	var userRequest request.UserUpdateRequest

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

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Init UpdateUser controller", err,
			zap.String("journey", "UpdateUser"),
		)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.service.UpdateUserService(userIdInt, domain)
	if err != nil {
		logger.Error("Error trying to call UpdateUser service", err,
			zap.String("journey", "UpdateUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User updated successfully",
		zap.String("journey", "UpdateUser"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))

}
