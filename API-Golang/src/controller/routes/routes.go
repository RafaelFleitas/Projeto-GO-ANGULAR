package routes

import (
	"github.com/RafaelFleitas/API-Golang/src/controller"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"github.com/gin-gonic/gin"
)

// Inicialização das rotas
func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", model.MiddlewareVerifyToken, userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", model.MiddlewareVerifyToken, userController.FindUserByEmail)
	r.GET("/getAllUsers", model.MiddlewareVerifyToken, userController.FindAllUsers)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", model.MiddlewareVerifyToken, userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", model.MiddlewareVerifyToken, userController.DeleteUser)

	r.POST("/login", userController.LoginUser)

}
