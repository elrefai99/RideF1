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

	db, err := config.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	defer sqlDB.Close()
	workerQueue := Queue.LeoxyWorker(10, db)
	defer workerQueue.Close()

	server.NewRouters(app, db, workerQueue)

	return server.NewServer(app)
}
