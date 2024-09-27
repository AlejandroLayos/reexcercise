package main

import (
	"context"
	"log"
	"time"

	"repartners/internal/adapter/config"
	"repartners/internal/adapter/handler/http"
	"repartners/internal/adapter/storage/redis"
	"repartners/internal/core/port"
	"repartners/internal/core/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title RE PARTNERS API
// @version 1.0
// @description This server implements a simple API for a PARTNERS test.

// @host localhost:8080
// @BasePath /packs
const (
	baseUserPath = "/packs"
)

func main() {
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4173"}, // Change this to your frontend's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Swagger handler
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	cache, err := redis.New(context.Background(), &config.Redis{Addr: "localhost:6379", Password: ""})
	if err != nil {
		log.Fatalf("Error creating Redis client: %v", err)
	}

	port.SetPackageServiceInstance(&service.PackServiceImpl{})
	port.SetConfigServiceInstance(&service.ConfigServiceImpl{CacheRepository: cache})

	packageHandler := http.NewPackageHandler(port.GetPackageServiceInstance(), port.GetConfigServiceInstance())

	r.GET(baseUserPath+"/:items", packageHandler.GetItemsHandler)
	r.PUT(baseUserPath+"/config", packageHandler.UpdateConfigHandler)

	r.Run(":8080")
}
