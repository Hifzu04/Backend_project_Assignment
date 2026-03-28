package main

import (
	controllers "ecommerce-backend/Controller"
	middleware "ecommerce-backend/Middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // React URL
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

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
	router.Use(CORSMiddleware())

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
