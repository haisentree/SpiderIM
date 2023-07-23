package rpcMsgGateway

import (
	"context"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	pbMsgGateway "SpiderIM/pkg/proto/msg_gateway"
	"SpiderIM/pkg/rabbitmq"
)

var (
	MessageProducer *rabbitmq.Producer
)

type rpcMsgGateway struct {
	rpcPort int
}

func New_rpcMsgGatewaySrv(port int) *rpcMsgGateway {
	return &rpcMsgGateway{
		rpcPort: 8849,
	}
}
func (rpc *rpcMsgGateway) RpcMsgGateway_Init() {
	MessageProducer = &rabbitmq.Producer{}
	MessageProducer.Producer_Init("test", "send")
}

func (rpc *rpcMsgGateway) ReceiveMessage(_ context.Context, req *pbMsgGateway.MessageReq) (*pbMsgGateway.MessageResp, error) {
	// zap.S().Info("msg:", req.Message)
	// zap.S().Info("type:", req.Type)
	log.Println("tpc client")

	return &pbMsgGateway.MessageResp{ErrCode: 200, ErrMsg: "error"}, nil
}

func (rpc *rpcMsgGateway) ReceiveClientMsg(_ context.Context, req *pbMsgGateway.ClientMsgReq) (*pbMsgGateway.MessageResp, error) {
	// zap.S().Info("msg:", req.Message)
	// zap.S().Info("type:", req.Type)
	log.Println("receiveClientMsg:", req.Content)
	MessageProducer.SendMessage([]byte(req.Content))
	// 这里接收的消息需要存储在mq中
	return &pbMsgGateway.MessageResp{ErrCode: 200, ErrMsg: "success"}, nil
}

func (rpc *rpcMsgGateway) Run() {

	address := "0.0.0.0" + ":" + strconv.Itoa(rpc.rpcPort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// pbMsgGateway.RegisterMsgGatewayServer(s, &rpcMsgGateway{})
	pbMsgGateway.RegisterMsgGatewayServer(s, &rpcMsgGateway{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
