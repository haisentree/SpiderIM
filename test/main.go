package main

import (
	// DBMysql "SpiderIM/pkg/db/mysql"
	// DBModel "SpiderIM/pkg/db/mysql/model"
	// "fmt"
	DBRedis "SpiderIM/pkg/db/redis"
	"fmt"
)

var (
	// MysqlDB DBMysql.MysqlDB
	RedisDB DBRedis.RedisDB
)

func main() {
	RedisDB.InitRedisDB()
	RedisDB.SetClientStatus(12, false)
	fmt.Println(RedisDB.GetClientStauts(12))
	// MysqlDB.InitMysqlDB()
	// MysqlDB.DB.AutoMigrate(&DBModel.Client{}, &DBModel.ClientToMessage{}, &DBModel.ClientMessage{})
	// client := DBModel.NewClient()
	// client.CreateClient(MysqlDB.DB, 1)

	// client_to_message := DBModel.NewClientToMessage()
	// // client_to_message.CreateClientToMessage(MysqlDB.DB, 4, 2)
	// t := client_to_message.FindByClientIDAndRecvID(MysqlDB.DB, 4, 2)

	// client_message := DBModel.NewClientMessage()
	// client_message.CreateMessage(MysqlDB.DB, t, 8, "hello8", true)
	// // result := client_message.FindMessageBySeq(MysqlDB.DB, 3, 6)
	// // for i, v := range result {
	// // 	fmt.Println(i, ":", v.Content)
	// // }

	// fmt.Println("end:", t)
}
