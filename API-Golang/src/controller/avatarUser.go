package controller

import (
	"fmt"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UploadAvatar(c *gin.Context) {

	logger.Info("Init UploadAvatar controller",
		zap.String("journey", "UploadAvatar"),
	)

	// Pega o ID do usuário logado, que o middleware colocou no contexto
	userIdValue, exists := c.Get("userId")
	if !exists {
		errRest := rest_err.NewUnauthorizedRequestError("Usuário não autenticado")
		c.JSON(errRest.Code, errRest)
		return
	}
	userId := userIdValue.(int64)

	file, err := c.FormFile("avatar")
	if err != nil {
		logger.Error("Error trying to upload avatar", err,
			zap.String("journey", "UploadAvatar"),
		)
		errRest := rest_err.NewBadRequestError("Arquivo não encontrado")
		c.JSON(errRest.Code, errRest)
		return
	}

	filePath := fmt.Sprintf("./uploads/user-avatars/%s", file.Filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		logger.Error("Error trying to save avatar", err,
			zap.String("journey", "UploadAvatar"),
		)
		errRest := rest_err.NewInternalServerError("Erro ao salvar arquivo")
		c.JSON(errRest.Code, errRest)
		return
	}

	avatarPath := fmt.Sprintf("/uploads/user-avatars/%s", file.Filename)

	// Salva a URL no banco, associada ao usuário logado
	updatedUser, restErr := uc.service.UpdateAvatarService(userId, avatarPath)
	if restErr != nil {
		logger.Error("Error trying to update avatar in database", restErr,
			zap.String("journey", "UploadAvatar"),
		)
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("Avatar updated successfully",
		zap.String("journey", "UploadAvatar"),
	)

	c.JSON(200, view.ConvertDomainToResponse(updatedUser))
}
