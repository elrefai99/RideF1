package main

import (
	"log"
	"os"

	"github.com/elrefai99/RideF1/app/api/internal/Common/database"
	"github.com/elrefai99/RideF1/app/api/internal/Queue"
	"github.com/elrefai99/RideF1/app/api/internal/middleware"
	"github.com/elrefai99/RideF1/app/api/internal/server"
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
	workerQueue := Queue.LeoxyWorker(10)
	defer workerQueue.Close()
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

	server.NewRouters(app, db, workerQueue)

	return server.NewServer(app)
}
