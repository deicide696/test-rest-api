package main

// @title			Bacu Test REST API
// @version		0.0.1

// @host			localhost:8080
// @BasePath	/api/test/v1

import (
	"log"
	"net/http"

	"github.com/deicide696/test-rest-api/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs" // TODO: Remove when add documentation
)

var Config = config.Config

func server() (*gin.Engine, *gin.RouterGroup) {
	s := gin.New()

	logger := gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/api/last-mile/v1/healthcheck"},
	})

	// TODO: We should review cors.Default()
	s.Use(logger, gin.Recovery(), cors.Default())

	router := s.Group("/api/last-mile/v1")

	if Config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	if Config.Environment != "production" {
		router.GET(
			"/docs/*any",
			ginSwagger.WrapHandler(
				swaggerFiles.Handler,
				ginSwagger.InstanceName(docs.SwaggerInfo.InstanceName()),
			),
		)
	}

	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	return s, router
}

func main() {
	serv, _ := server()

	if err := serv.Run(); err != nil {
		log.Fatalf("fatal error: %s", err.Error())
	}
}
