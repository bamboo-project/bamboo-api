package user

import (
	"net/http"

	"bamboo-api/app/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUser(c *gin.Context) {
	walletId := c.GetHeader("walletId")
	if walletId == "" {
		log.Warn("[GetUser] the user id params is nil")
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	resp, err := service.UserService.GetUser(walletId)
	if nil != err {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
