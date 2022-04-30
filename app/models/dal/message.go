package dal

import (
	"bamboo-api/app/pkg/entity/dto"
	"bamboo-api/app/pkg/entity/po"
)

func CreateMessage(createMessage *po.CreateMessage) *dto.Response {
	return &dto.Response{Data: "1233"}
}

func QueryMessagePage(toWalletId string) *dto.Response {
	wrapper := make(map[string]interface{}, 0)
	wrapper["toWalletId"] = toWalletId
	//messageList := database.GetDBClient().Model(dao.Message{}).Where(wrapper)
	return nil
}
