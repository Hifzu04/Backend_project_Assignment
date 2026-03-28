package main

import (
	controllers "ecommerce-backend/Controller"
	middleware "ecommerce-backend/Middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Make sure it exists in the root directory.")
	}

	mongoURI := os.Getenv("MONGO_URL")
	if mongoURI == "" {
		log.Fatal("MONGO_URL is not set in the .env file")
	}
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	// API Versioning
	v1 := router.Group("/api/v1")
	{
		// Public Auth Routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		// Protected Product Routes
		products := v1.Group("/products")
		products.Use(middleware.Authenticate())
		{
			products.GET("/", controllers.GetProducts)
			// Only Admin can add products
			products.POST("/", middleware.OnlyAdmin(), controllers.CreateProduct)
		}
	}

	log.Printf("Server starting on port %s", port)
	router.Run(":" + port)
}
