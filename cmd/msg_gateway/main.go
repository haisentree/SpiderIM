package main

import (
	MsgGateway "SpiderIM/internal/msg_gateway"
)

func main() {
	var myWSGate MsgGateway.WServer

	myWSGate.OnInit(8848)
	myWSGate.Run()
}
