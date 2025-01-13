package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nedpals/supabase-go"
	"Waste-Management/Auth"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize Supabase client
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	if supabaseUrl == "" || supabaseKey == "" {
		log.Fatal("Supabase URL or key is missing in .env")
	}

	client := supabase.CreateClient(supabaseUrl, supabaseKey)
	if client == nil {
		log.Fatal("Failed to initialize Supabase client")
	}

	// Initialize auth services
	authService := Auth.NewAuthService(client)

	// Set Gin mode to release
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// OPTIONS handler
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.AbortWithStatus(204)
	})

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, Gin!"})
	})

	r.POST("/signup", authService.Signup)
	r.POST("/login", authService.Login)

	// Start server
	fmt.Println("Starting server on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
