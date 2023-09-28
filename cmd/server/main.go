package main

import (
	"Wellness-monitoring/config"
	"Wellness-monitoring/controllers"
	"Wellness-monitoring/models"
	"Wellness-monitoring/routes"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var (
	server         *gin.Engine
	ctx            context.Context
	DB             *gorm.DB
	AuthController controllers.AuthController
	AuthRoutes     routes.AuthRoutes
)

func init() {
	ctx = context.TODO()
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Create a new Dialector for PostgreSQL
	dialector := postgres.New(postgres.Config{
		DSN: config.PostgresSource,
	})

	// Use the dialector to open a connection
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}

	//creating table in postgres
	db.AutoMigrate(&models.User{})
	fmt.Println("PostgreSQL connected successfully...")

	AuthController = *controllers.NewAuthController(db, ctx)
	AuthRoutes = routes.NewAuthRoutes(AuthController, db)

	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Origin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")

	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Welcome to Golang with PostgreSQL"})
	})

	AuthRoutes.AuthRoute(router)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": fmt.Sprintf("Route %s not found", ctx.Request.URL)})
	})
	log.Fatal(server.Run(":" + config.Port))
}
