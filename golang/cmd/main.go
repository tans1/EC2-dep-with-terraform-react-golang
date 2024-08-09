package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	v1routes "github.com/tans1/go-web-server/api/v1/routes"
	"github.com/tans1/go-web-server/config"
)

func main() {
	file, err := os.OpenFile("logs/errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	err = godotenv.Load()
	if err != nil {
		log.Printf("%s:Error loading .env file ", "main.go")
	}

	// db configuration
	config := config.NewDbConfig()
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatalf("%s:Unable to connect to db", "main.go")
	}

	// server
	gin.SetMode(gin.DebugMode) // running mode, for production we change it to gin.ReleaseMode
	server := gin.Default()
	server.Use(cors.New(cors.Config{
        AllowOrigins: []string{
            "https://foo.com",
            "http://localhost:5173", // Allowing localhost with port 5173
            "http://localhost:8080", // Allowing localhost with port 8080
        },
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))
	//   cors.Default()

	// v1 is a version 1.0 routes
	v1 := server.Group("v1")
	v1routes.RegisterRoutes(v1, db)

	port := os.Getenv("PORT")
	fmt.Print("Starting server on port: ", port)
	server.Run(fmt.Sprintf(":%s", port))
}
