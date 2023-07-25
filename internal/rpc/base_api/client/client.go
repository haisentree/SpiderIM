package rpcBaseAPIClient

import (
	pbBaseAPIClient "SpiderIM/pkg/proto/base_api/client"
	"context"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type rpcBaseAPIClient struct {
	rpcPort int
}

func New_rpcBaseAPIClient(port int) *rpcBaseAPIClient {
	return &rpcBaseAPIClient{
		rpcPort: 8850,
	}
}

func (rpc *rpcBaseAPIClient) CreateClient(_ context.Context, req *pbBaseAPIClient.CreateMessageReq) (*pbBaseAPIClient.CreateMessageResp, error) {
	// zap.S().Info("msg:", req.Message)
	// zap.S().Info("type:", req.Type)
	log.Println("tpc client")

	return &pbBaseAPIClient.CreateMessageResp{ClientID: "23", ClientUUID: "sdfgdasfsdaf-asfsfadf-sdfsd"}, nil
}

func (rpc *rpcBaseAPIClient) Run() {

	address := "0.0.0.0" + ":" + strconv.Itoa(rpc.rpcPort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// pbMsgGateway.RegisterMsgGatewayServer(s, &rpcMsgGateway{})
	pbBaseAPIClient.RegisterMsgGatewayServer(s, &rpcBaseAPIClient{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
