package ping

import (
	"net/http"

	"bamboo-api/app/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func PingHandler(c *gin.Context) {
	log.WithFields(log.Fields{
		"func":        "PingHandler",
		"status_code": http.StatusOK,
	}).Info("success")

	c.JSON(http.StatusOK, utils.Response{Data: "ping", StatusCode: http.StatusOK})
}
