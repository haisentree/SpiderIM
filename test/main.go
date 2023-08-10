package main

import (
	DBMysql "SpiderIM/pkg/db/mysql"
	DBModel "SpiderIM/pkg/db/mysql/model"
	"fmt"
)

var (
	MysqlDB DBMysql.MysqlDB
)

func main() {
	MysqlDB.InitMysqlDB()
	MysqlDB.DB.AutoMigrate(&DBModel.Client{}, &DBModel.ClientToMessage{},&DBModel.ClientMessage{})
	client := DBModel.NewClient()
	client.CreateClient(MysqlDB.DB, 1)

	client_to_message := DBModel.NewClientToMessage()
	client_to_message.CreateClientToMessage(MysqlDB.DB, 4, 2)
	t := client_to_message.FindByClientIDAndRecvID(MysqlDB.DB, 4, 2)

	client_message := DBModel.NewClientMessage()
	client_message.CreateMessage(MysqlDB.DB,t,2,"hello")

	fmt.Println("end:", t)
}
