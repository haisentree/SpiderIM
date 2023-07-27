package rpcBaseAPIClient

import "gorm.io/gorm"

type ClientModel struct {
	gorm.Model
	ClientUUID string `gorm:unique`
	ClientType int32
}

func (ClientModel) TableName() string {
	return "client"
}
