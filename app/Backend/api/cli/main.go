package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/elrefai99/PBE/app/Backend/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	app := gin.Default()
	app.Use(gin.Logger())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	_, err := config.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{Addr: os.Getenv("PORT"), Handler: app}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is Run")
}
