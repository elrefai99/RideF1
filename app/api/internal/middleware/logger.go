package middleware

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()
		req_id, _ := c.Get("req_Id")

		log.Printf(
			"%s %s %s %d %s",
			req_id,
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			time.Since(start),
		)
		Log.Info(
			"HTTP Request",
			"request_id", c.GetString("req_Id"),
			"Method", c.Request.Method,
			"Path", c.Request.URL.Path,
			"Status", c.Writer.Status(),
			"Time", time.Since(start),
			"Size", c.Writer.Size(),
		)
	}
}

var Log *slog.Logger

func Init() error {
	if err := os.MkdirAll("logs", 0755); err != nil {
		return err
	}

	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	Log = slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	return nil
}
