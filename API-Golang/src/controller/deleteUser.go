package controller

import (
	"net/http"
	"strconv"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// DeleteUser godoc
// @Summary Remove um usuário
// @Description Exclui um usuário existente pelo ID
// @Tags users
// @Produce json
// @Param userId path int true "ID do usuário"
// @Security BearerAuth
// @Success 200
// @Failure 400 {object} rest_err.RestErr
// @Failure 404 {object} rest_err.RestErr
// @Router /deleteUser/{userId} [delete]
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
