package message

import (
	"net/http"

	"bamboo-api/app/models/dal"
	"bamboo-api/app/pkg/entity/po"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateMessage(c *gin.Context) {
	content := c.Param("content")
	if content == "" {
		log.Warn("")
	}
	createMsg := &po.CreateMessage{
		Content: content,
	}
	c.JSON(http.StatusOK, dal.CreateMessage(createMsg))
}

func GetMessages(c *gin.Context) {
	content := c.Param("content")
	if content == "" {
		log.Warn("")
	}
	createMsg := &po.CreateMessage{
		Content: content,
	}
	c.JSON(http.StatusOK, dal.CreateMessage(createMsg))
}
