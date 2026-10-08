package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type demoEmail struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
}

type sendEmailRequest struct {
	From     string   `json:"from" binding:"required,email"`
	To       []string `json:"to" binding:"required,min=1,dive,email"`
	Subject  string   `json:"subject" binding:"required,max=200"`
	Message  string   `json:"message" binding:"required"`
	Provider string   `json:"provider" binding:"required,oneof=gmail resend"`
}

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), localCors())

	router.GET("/api/emails/demo", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, demoEmail{
			From:    "hello@pbe.com",
			To:      []string{"maya@example.com"},
			Subject: "A little note for your day",
			Message: "Hi Maya,\n\nJust wanted to send a little note and let you know we're thinking of you. Hope your day is going well!\n\nWarmly,\nThe PBE team",
		})
	})

	router.POST("/api/emails/send", func(ctx *gin.Context) {
		var request sendEmailRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid email request"})
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{
			"status":          "accepted",
			"provider":        request.Provider,
			"recipient_count": len(request.To),
			"message":         "Demo email accepted. No email was delivered.",
		})
	})

	return router
}

func localCors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		origin := ctx.GetHeader("Origin")
		if origin == "http://localhost:3000" || origin == "http://127.0.0.1:3000" {
			ctx.Header("Access-Control-Allow-Origin", origin)
			ctx.Header("Vary", "Origin")
			ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			ctx.Header("Access-Control-Allow-Headers", "Content-Type")
		}

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.Next()
	}
}
