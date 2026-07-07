package controller

import (
	"fmt"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UploadAvatar(c *gin.Context) {

	logger.Info("Init UploadAvatar controller",
		zap.String("journey", "UploadAvatar"),
	)

	file, err := c.FormFile("avatar")

	if err != nil {
		logger.Error("Error trying to upload avatar", err,
			zap.String("journey", "UploadAvatar"),
		)
		return
	}

	filePath := fmt.Sprintf("./uploads/user-avatars/%s", file.Filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		logger.Error("Error trying to save avatar", err,
			zap.String("journey", "UploadAvatar"),
		)
		return
	}

	imageURL := fmt.Sprintf("http://localhost:8000/uploads/user-avatars/%s", file.Filename)

	c.JSON(200, gin.H{
		"message": "Arquivo enviado com sucesso",
		"url":     imageURL,
	})
}
