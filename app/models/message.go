package models

import "time"

type Message struct {
	id                  int64 `gorm:"-;primary_key;AUTO_INCREMENT"`
	FromWallet          string
	FromWalletType      string
	FromWalletAvatarURL string
	ToWallet            string
	ToWalletType        string
	ToWalletAvatarURL   string
	CreateAt            time.Time
	Content             string
}

func (m *Message) TableName() string {
	return "bamboo.message"
}
