package utils

import (
	"github.com/bytedance/sonic"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

func (response *Response) ResponseToJSON() string {
	resp, _ := sonic.MarshalString(response)
	return resp
}
