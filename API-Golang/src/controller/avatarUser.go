package controller

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model"
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

	ext := filepath.Ext(file.Filename) // ex: ".jpg"
	avatarFileName := model.HashFileName(fmt.Sprintf("%d", userId)) + ext

	hash := model.HashFileName(fmt.Sprintf("%d", userId))
	avatarFileName = hash + ext

	oldFiles, _ := filepath.Glob(fmt.Sprintf("%s%s*", model.AvatarUploadDir, hash))

	for _, oldFile := range oldFiles {
		if err := os.Remove(oldFile); err != nil {
			logger.Error("Error trying to remove old avatar", err,
				zap.String("journey", "UploadAvatar"),
			)
		}
	}

	filePath := fmt.Sprintf("%s%s", model.AvatarUploadDir, avatarFileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		logger.Error("Error trying to save avatar", err,
			zap.String("journey", "UploadAvatar"),
		)
		errRest := rest_err.NewInternalServerError("Erro ao salvar arquivo")
		c.JSON(errRest.Code, errRest)
		return
	}

	// Salva a URL no banco, associada ao usuário logado

	updatedUser, restErr := uc.service.UpdateAvatarService(userId, avatarFileName)

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
