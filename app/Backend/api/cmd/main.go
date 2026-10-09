package main

import (
	"log"
	"os"

	"github.com/elrefai99/PBE/app/Backend/api/internal/database"
	"github.com/elrefai99/PBE/app/Backend/api/internal/middleware"
	"github.com/elrefai99/PBE/app/Backend/api/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	gin.SetMode(gin.ReleaseMode)

	app := gin.Default()

	db, err := database.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	defer sqlDB.Close()

	app.Use(
		middleware.Logger(),
		middleware.Recovery(),
		middleware.CORS("http://localhost:3000"),
	)

	server.NewRouters(app)

	return server.NewServer(app)
}
