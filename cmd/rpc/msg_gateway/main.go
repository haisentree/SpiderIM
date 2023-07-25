package main

import (
	rpcMsgGateway "SpiderIM/internal/rpc/msg_gateway"
	"fmt"
)

func main() {
	rpcServer := rpcMsgGateway.New_rpcMsgGatewaySrv(8849)
	rpcServer.RpcMsgGateway_Init()
	fmt.Println("running")
	rpcServer.Run()
}
