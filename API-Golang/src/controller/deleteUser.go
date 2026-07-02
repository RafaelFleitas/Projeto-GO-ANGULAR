package controller

import (
	"net/http"
	"strconv"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Controller do Delete
func (uc *userControllerInterface) DeleteUser(c *gin.Context) {

	logger.Info("Init deleteUser controller",
		zap.String("journey", "deleteUser"),
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

	err := uc.service.DeleteUserService(userIdInt)

	if err != nil {
		logger.Error("Error trying to call deleteUser service", err,
			zap.String("journey", "deleteUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User deleted successfully",
		zap.String("journey", "deleteUser"),
	)

	c.Status(http.StatusOK)

}
