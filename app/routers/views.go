package routers

import (
	"bamboo-api/app/clients/database"
	"bamboo-api/app/dao"
	"bamboo-api/app/routers/message"
	"bamboo-api/app/routers/ping"
	"bamboo-api/app/routers/twitter"
	"bamboo-api/app/routers/user"
	"bamboo-api/app/service"

	"github.com/gin-gonic/gin"
)

var (
	MessageService *service.MessageService
)

func InitRouters(e *gin.Engine) {
	// init dao
	msgDao := dao.NewMessageDao(database.GetDBClient())
	userDao := dao.NewUserDao(database.GetDBClient())
	// init services
	service.InitMessageService(msgDao)
	service.InitUserService(userDao)
	initPing(e.Group("/"))
	createUserMessage(e.Group("/api"))
	messageRouter(e.Group("/api"))
	twitterRoute(e.Group("/api"))
	userRoute(e.Group("/api"))
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

func twitterRoute(r *gin.RouterGroup) {
	r.GET("/login/twitter/auth", twitter.LoginByTwitter)
	r.GET("/login/twitter/auth/callback", twitter.TwitterCallback)
}
func userRoute(r *gin.RouterGroup) {
	r.GET("/user", user.GetUser)
}
