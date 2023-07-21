package MsgGateway

import (
	pbMsgGateway "SpiderIM/pkg/proto/msg_gateway"
	"log"

	"google.golang.org/grpc"
)

var (
	MsgGatewaySrvClient pbMsgGateway.MsgGatewayClient
)

func init() {
	SrvClient_Init()
}

func SrvClient_Init() {
	conn, err := grpc.Dial("0.0.0.0:8849", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	msgGatewaySrvClient := pbMsgGateway.NewMsgGatewayClient(conn)
	MsgGatewaySrvClient = msgGatewaySrvClient
}
