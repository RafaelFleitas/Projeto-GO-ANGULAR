package controller

import (
	"net/http"
	"net/mail"
	"strconv"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/controller/model/response"
	"github.com/RafaelFleitas/API-Golang/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// FindUserByID godoc
// @Summary Busca um usuário pelo ID
// @Description Retorna os dados de um usuário específico
// @Tags users
// @Produce json
// @Param userId path int true "ID do usuário"
// @Security BearerAuth
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 404 {object} rest_err.RestErr
// @Router /getUserById/{userId} [get]
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

// FindUserByEmail godoc
// @Summary Busca um usuário pelo email
// @Description Retorna os dados de um usuário específico a partir do email
// @Tags users
// @Produce json
// @Param userEmail path string true "Email do usuário"
// @Security BearerAuth
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 404 {object} rest_err.RestErr
// @Router /getUserByEmail/{userEmail} [get]
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

// FindAllUsers godoc
// @Summary Lista todos os usuários
// @Description Retorna a lista paginada de usuários cadastrados
// @Tags users
// @Produce json
// @Param page query int false "Número da página (padrão: 1)"
// @Param pageSize query int false "Itens por página (padrão: 10)"
// @Security BearerAuth
// @Success 200 {object} response.PaginatedUsersResponse
// @Failure 401 {object} rest_err.RestErr
// @Router /getAllUsers [get]
func (uc *userControllerInterface) FindAllUsers(c *gin.Context) {
	logger.Info("Init FindAllUsersController",
		zap.String("journey", "FindAllUsers"),
	)

	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.ParseInt(c.DefaultQuery("pageSize", "10"), 10, 64)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	search := c.Query("search")

	usersDomain, total, restErr := uc.service.FindAllUsersService(page, pageSize, search)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	totalPages := (total + pageSize - 1) / pageSize

	logger.Info("FindAllUsersController successfully executed",
		zap.String("journey", "FindAllUsers"),
	)

	c.JSON(http.StatusOK, response.PaginatedUsersResponse{
		Users:      view.ConvertDomainListToResponse(usersDomain),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	})
}
