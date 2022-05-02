package message

import (
	"net/http"
	"strconv"
	"time"

	"bamboo-api/app/models"
	"bamboo-api/app/pkg/entity/dto"
	"bamboo-api/app/service"

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

func GetMessageList(c *gin.Context) {
	toAddress := c.DefaultQuery("toAddress", "")
	count := c.DefaultQuery("count", "5")
	cnt, err := strconv.Atoi(count)

	if toAddress == "" || err != nil {
		c.JSON(http.StatusBadRequest, &dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "params invalid",
		})
		return
	}

	msgs := service.MsgService.GetMessageList(toAddress, int64(cnt))

	c.JSON(http.StatusOK, &dto.Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: map[string]interface{}{
			"msg_len":  len(msgs),
			"msg_list": msgs,
		},
	})

}

func PostMessage(c *gin.Context) {
	content := c.DefaultPostForm("content", "")
	fromAddress := c.DefaultPostForm("fromAddress", "")
	toAddress := c.DefaultPostForm("toAddress", "")
	fromAddressAvatarURL := c.DefaultPostForm("fromAddressAvatarURL", "")
	toAddressAvatarURL := c.DefaultPostForm("toAddressAvatarURL", "")
	if content == "" || fromAddress == "" || toAddress == "" {
		c.JSON(http.StatusBadRequest, &dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "params invalidÏ",
		})
		return
	}

	msg := models.Message{
		FromWallet:          fromAddress,
		ToWallet:            toAddress,
		Content:             content,
		FromWalletAvatarURL: fromAddressAvatarURL,
		ToWalletAvatarURL:   toAddressAvatarURL,
		CreateAt:            time.Now(),
	}
	err := service.MsgService.PostNewMessage(&msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "server internal err",
		})
		return
	}
	c.JSON(http.StatusOK, &dto.Response{
		Code: http.StatusOK,
		Msg:  "success",
	})
}
