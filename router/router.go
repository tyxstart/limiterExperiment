package router

import (
	"github.com/gin-gonic/gin"
	v1 "limiterExperiment/api/v1"
	"limiterExperiment/config"
)

func InitRouter() {
	gin.SetMode(config.AppMode)

	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("test", v1.Test)
		router.POST("setLimiter", v1.SetLimiter)
	}

	r.Run(config.HttpPort)
}
