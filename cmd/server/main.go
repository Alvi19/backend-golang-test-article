package main

import (
	"fmt"
	"log"

	_ "github.com/Alvi19/backend-golang-test/docs"
	"github.com/Alvi19/backend-golang-test/internal/config"
	"github.com/Alvi19/backend-golang-test/internal/delivery/router"
	"github.com/Alvi19/backend-golang-test/internal/repository"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Article API
// @version 1.0
// @description API documentation for Article App (Echo + Golang).
// @host localhost:8080
// @BasePath /api
func main() {
	cfg, err := config.NewConfigFromEnv()
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}

	db, err := repository.NewPostgresGorm(cfg)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	if err := repository.AutoMigrate(db); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	e := router.NewRouter(db, cfg)

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	// Swagger route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("[Article API] starting Echo server on %s | env=%s | db=%s",
		addr, cfg.AppEnv, cfg.DBName)

	e.Logger.Fatal(e.Start(addr))
}
