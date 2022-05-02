package models

import "time"

type Message struct {
	ID                  uint      `json:"id" gorm:"primarykey"`
	FromWallet          string    `json:"from_wallet"`
	FromWalletType      string    `json:"from_wallet_type"`
	FromWalletAvatarURL string    `json:"from_wallet_avatar_url"`
	ToWallet            string    `json:"to_wallet"`
	ToWalletType        string    `json:"to_wallet_type"`
	ToWalletAvatarURL   string    `json:"to_wallet_avatar_url"`
	CreateAt            time.Time `json:"create_at"`
	Content             string    `json:"content"`
}

func (m *Message) TableName() string {
	return "bamboo.message"
}
