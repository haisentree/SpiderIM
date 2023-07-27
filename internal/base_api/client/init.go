package baseAPIClient

import (
	pbBaseAPIClient "SpiderIM/pkg/proto/base_api/client"
	"log"

	"google.golang.org/grpc"
)

var (
	BaseAPIClientSrvClient pbBaseAPIClient.BaseAPIClientClient
)

func init() {
	SrvClient_Init()
}

func SrvClient_Init() {
	conn, err := grpc.Dial("0.0.0.0:8851", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	baseAPIClientSrvClient := pbBaseAPIClient.NewBaseAPIClientClient(conn)
	BaseAPIClientSrvClient = baseAPIClientSrvClient
}
