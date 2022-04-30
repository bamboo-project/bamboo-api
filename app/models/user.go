package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	WalletAddress string
	WalletType    string
	BambooCoins   int64
	AvatarUrl     string
	Twitter       string
}

func (u *User) TableName() string {
	return "bamboo.users"
}
