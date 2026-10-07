package main

import (
	"log"
	"net/http"

	"github.com/elrefai99/PBE/app/Backend/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	env, err := config.DotenvConfig()
	if err != nil {
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

	server := &http.Server{Addr: env.PORT, Handler: app}
	server.ListenAndServe()
}
