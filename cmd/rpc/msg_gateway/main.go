package main

import (
	"SpiderIM/internal/rpc/msg_gateway"
	"fmt"
)

func main() {
	rpcServer := rpcMsgGateway.New_rpcMsgGatewaySrv(8849)
	fmt.Println("running")
	rpcServer.Run()
}