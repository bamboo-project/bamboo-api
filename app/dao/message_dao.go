package dao

import (
	"bamboo-api/app/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MessageDao struct {
	db *gorm.DB
}

func NewMessageDao(db *gorm.DB) *MessageDao {
	return &MessageDao{db: db}
}

func (m *MessageDao) Create(msg *models.Message) error {
	err := m.db.Create(msg).Error
	if err != nil {
		log.Warnf("[MessageDao] create err:%+v", err)
		return err
	}
	return nil
}

func (m *MessageDao) Query(toWallet string, count int64) ([]*models.Message, error) {
	msgs := make([]*models.Message, 0)
	err := m.db.Where("to_wallet = ?", toWallet).Order("create_at desc").Limit(int(count)).Find(&msgs).Error
	if err != nil {
		log.Warnf("[MessageDao] query err, key: %+v, err:%+v", toWallet, err)
		return nil, err
	}
	return msgs, nil
}
