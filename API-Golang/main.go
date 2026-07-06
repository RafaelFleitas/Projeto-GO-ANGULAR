package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/RafaelFleitas/API-Golang/docs"
	oraclesql "github.com/RafaelFleitas/API-Golang/src/configuration/database/oracleSQL"
	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func corsMiddleware() gin.HandlerFunc {
	allowedOrigins := map[string]bool{}
	rawOrigins := os.Getenv("ALLOWED_ORIGINS")

	if rawOrigins == "" {
		rawOrigins = "http://localhost:4200"
	}

	for _, origin := range strings.Split(rawOrigins, ",") {
		allowedOrigins[strings.TrimSpace(origin)] = true
	}

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		c.Writer.Header().Set("Vary", "Origin")
		if allowedOrigins[origin] {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// @title API CRUD Golang
// @version 1.0
// @description API de gerenciamento de usuários com autenticação JWT.
// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {

	logger.Info("About to start the application")
	godotenv.Load()

	//Inicializa o banco de dados ORACLE
	db, err := oraclesql.NewOracleConnection()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	userController := initDependencies(db)

	router := gin.Default()
	router.Use(corsMiddleware())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}

}
