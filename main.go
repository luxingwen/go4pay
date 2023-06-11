package main

import (
	"go4pay/controller"
	log "go4pay/pkg/logger"
	"go4pay/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.LoggerMiddleware, middleware.OpenFixAuth)
	r.POST("/pay/payWebhook/payCallback/openFix/*event", controller.OpenFixCallBack)

	err := r.Run(":9898")
	if err != nil {
		log.Errorf("Failed to start server: %v", err)
	}
}
