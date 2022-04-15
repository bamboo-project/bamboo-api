package router

import (
	"bamboo-api/app/router/ping"
	"github.com/gin-gonic/gin"
)

func InitRouters(e *gin.Engine) {
	initPing(e.Group("/"))
}

func initPing(r *gin.RouterGroup) {
	r.GET("/ping", ping.PingHandler)
}
