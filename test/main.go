package main

import (
	DBmongo "SpiderIM/pkg/db/mongodb"
)

var (
	MongoDB DBmongo.MongoDB
)

func main() {
	MongoDB.Init_mongodb()
	// MongoDB.Insert_client_msg()
	MongoDB.Update_client_msg()
	// MongoDB.Find_client_msg()
}
