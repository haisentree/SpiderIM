package main

import (
	rpcBaseAPIClient "SpiderIM/internal/rpc/base_api/client"
)

func main() {
	rpcServer := rpcBaseAPIClient.New_rpcBaseAPIClient(8851)
	rpcServer.RpcBaseAPIClient_Init()
	rpcServer.Run()
}
