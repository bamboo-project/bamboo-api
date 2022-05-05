package service

import (
	"errors"

	"bamboo-api/app/dao"
	"bamboo-api/app/models"
)

type MessageService struct {
	messageDao *dao.MessageDao
}

var MsgService *MessageService

func InitMessageService(messageDao *dao.MessageDao) {
	MsgService = &MessageService{messageDao: messageDao}
}

func (m *MessageService) PostNewMessage(msg *models.Message) error {
	if msg.FromWallet == msg.ToWallet {
		return errors.New("can't post to yourself")
	}
	//

	return m.messageDao.Create(msg)
}

func (m *MessageService) GetMessageList(toAddress string, count int64) []*models.Message {
	msgs, err := m.messageDao.Query(toAddress, count)
	if err != nil || msgs == nil {
		return nil
	}
	return msgs
}
