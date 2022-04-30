package routers

import (
	"bamboo-api/app/routers/message"
	"bamboo-api/app/routers/ping"

	"github.com/gin-gonic/gin"
)

func InitRouters(e *gin.Engine) {
	initPing(e.Group("/"))
	createUserMessage(e.Group("/api"))
}

func initPing(r *gin.RouterGroup) {
	r.GET("/ping", ping.PingHandler)
}

func createUserMessage(r *gin.RouterGroup) {
	r.POST("/user/message", message.CreateMessage)
}
