package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint           `json:"id" gorm:"primarykey"`
	CreatedAt     time.Time      `json:"create_at"`
	UpdatedAt     time.Time      `json:"update_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	WalletAddress string         `json:"wallet_address" gorm:"size:255"`
	WalletType    string         `json:"wallet_type"`
	BambooCoins   int64          `json:"bamboo_coins"`
	AvatarUrl     string         `json:"avatar_url"`
	TwitterUrl    string         `json:"twitter_url"`
	TwitterToken  string         `json:"twitter_token"`
	IsTwitter     uint8          `json:"is_twitter"`
	Followers     int64          `json:"followers"`
	Following     int64          `json:"following"`
}

func (u *User) TableName() string {
	return "bamboo.users"
}
