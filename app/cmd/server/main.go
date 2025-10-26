package main

import (
	"devops-portfolio/app/internal/api"
	"devops-portfolio/app/internal/cache"
	"devops-portfolio/app/internal/config"
	"devops-portfolio/app/internal/db"
	"devops-portfolio/app/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	cfg := config.Load()
	_ = godotenv.Load() // .env 있으면 로드
	pool := db.Connect(cfg)
	redisClient := cache.Connect(cfg)

	userService := service.NewUserService(pool, redisClient)
	handler := api.NewHandler(userService)

	r := gin.Default()
	r.GET("/", handler.RootCheck)
	r.GET("/healthz", handler.Health)
	r.POST("/users", handler.CreateUser)
	r.GET("/users/:id", handler.GetUser)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	log.Printf("Server running on port %s", cfg.AppPort)
	r.Run(":" + cfg.AppPort)
}
