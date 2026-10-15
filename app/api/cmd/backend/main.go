package main

import (
	"log"
	"os"

	"github.com/elrefai99/RideF1/app/api/internal/Common/config"
	"github.com/elrefai99/RideF1/app/api/internal/Queue"
	"github.com/elrefai99/RideF1/app/api/internal/middleware"
	"github.com/elrefai99/RideF1/app/api/internal/server"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if err := config.LoadEnv(); err != nil {
		return err
	}
	if err := middleware.Init(); err != nil {
		return err
	}

	gin.SetMode(gin.ReleaseMode)

	app := gin.Default()

	// Server Settings
	app.Use(
		middleware.Request_ID(),
		middleware.Logger(),
		middleware.Recovery(),
		middleware.CORS(
			os.Getenv("FRONTEND_DEV_MODE"),
			os.Getenv("FRONTEND_LIVE_MODE"),
			os.Getenv("SITE_DOMAIN"),
		),
	)

	// Database Connect
	db, err := config.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	redisClient, err := config.ConnectRedis(os.Getenv("REDIS_URL"))
	if err != nil {
		return err
	}
	defer redisClient.Close()

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	defer sqlDB.Close()

	// Worker Pool
	workerQueue := Queue.LeoxyWorker(10, db)
	defer workerQueue.Close()

	// Endpoints
	server.NewRouters(app, db, workerQueue)

	return server.NewServer(app)
}
