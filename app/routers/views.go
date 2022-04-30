package routers

import (
	"bamboo-api/app/clients/database"
	"bamboo-api/app/dao"
	"bamboo-api/app/routers/message"
	"bamboo-api/app/routers/ping"
	"bamboo-api/app/service"

	"github.com/gin-gonic/gin"
)

var (
	MessageService *service.MessageService
)

func InitRouters(e *gin.Engine) {

	// init dao
	msgDao := dao.NewMessageDao(database.GetDBClient())

	// init services
	service.InitMessageService(msgDao)

	initPing(e.Group("/"))
	createUserMessage(e.Group("/api"))
	messageRouter(e.Group("/api"))

}

func initPing(r *gin.RouterGroup) {
	r.GET("/ping", ping.PingHandler)
}

func createUserMessage(r *gin.RouterGroup) {
	r.POST("/user/message", message.CreateMessage)
}

func messageRouter(r *gin.RouterGroup) {

	r.POST("/user/message/v1", message.PostMessage)

	r.GET("/user/message/v1", message.GetMessageList)
}
