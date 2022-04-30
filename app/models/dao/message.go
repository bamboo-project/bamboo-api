package dao

import (
	"bamboo-api/app/pkg/entity/dto"
)

type Message struct {
	fromWallet     string
	fromWalletType string
	toWallet       string
	toWalletType   string
	createDate     int64
	content        string
}

func CreateMessage(message Message) *dto.Response {
	return &dto.Response{Data: "1233"}
}
