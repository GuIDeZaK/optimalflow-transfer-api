package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/guide-backend/internal/handler"
	"github.com/guide-backend/internal/model"
	"github.com/guide-backend/internal/repository"
	"github.com/guide-backend/internal/service"
	"github.com/guide-backend/pkg/jwt"
	"github.com/guide-backend/pkg/middleware"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env not found, continuing with system environment")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	// Migrate schema
	db.AutoMigrate(&model.User{})

	app := fiber.New()

	userRepo := repository.NewUserRepo(db)
	transferRepo := repository.NewTransferRepo(db)

	jwtService := jwt.NewJWTService()
	userService := service.NewUserService(userRepo, jwtService)
	transferService := service.NewTransferService(transferRepo, userRepo)

	userHandler := handler.NewUserHandler(userService)
	transferHandler := handler.NewTransferHandler(transferService)

	// Routes
	app.Post("/users", userHandler.CreateUser)
	app.Post("/login", userHandler.Login)
	app.Get("/users", userHandler.ListAllUsers)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Post("/transfer", middleware.AuthMiddleware(), transferHandler.Transfer)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
