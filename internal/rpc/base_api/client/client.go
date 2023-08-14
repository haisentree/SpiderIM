package rpcBaseAPIClient

import (
	DBMysql "SpiderIM/pkg/db/mysql"
	DBModel "SpiderIM/pkg/db/mysql/model"
	pbBaseAPIClient "SpiderIM/pkg/proto/base_api/client"
	pkgMessage "SpiderIM/pkg/public/message"
	"context"
	"log"
	"net"
	"strconv"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

var (
	MysqlDB DBMysql.MysqlDB
)

type rpcBaseAPIClient struct {
	rpcPort int
}

func New_rpcBaseAPIClient(port int) *rpcBaseAPIClient {
	return &rpcBaseAPIClient{
		rpcPort: 8851,
	}
}

func (rpc *rpcBaseAPIClient) RpcBaseAPIClient_Init() {
	MysqlDB.InitMysqlDB()
	MysqlDB.DB.AutoMigrate(&DBModel.Client{})
}

// client,group,server

func (rpc *rpcBaseAPIClient) CreateClient(_ context.Context, req *pbBaseAPIClient.CreateClientReq) (*pbBaseAPIClient.CreateClientResp, error) {
	// 从配置文件中读取
	var WsClientCreateKey = "dsfrserererst"
	var WsServerCreateKey = "sdfdsfdsfsd"

	// 1.校验key
	// 如果处理错误，直接return nil，error，不用return空结构体数据
	switch req.ClientType {
	case pkgMessage.Internal_Platform:
		log.Println("deal server key")
		if req.SecretKey != WsServerCreateKey {
			return &pbBaseAPIClient.CreateClientResp{}, grpc.Errorf(400, "Server Secret Key Fail")
		}
	case pkgMessage.Android_Platform, pkgMessage.IOS_Platform, pkgMessage.PC_Platform, pkgMessage.Web_Platform:
		log.Println("deal client key")
		if req.SecretKey != WsClientCreateKey {
			return &pbBaseAPIClient.CreateClientResp{}, grpc.Errorf(400, "Client Secret Key Fail")
		}
	default:
		log.Println("Secret Key invalid")
		return &pbBaseAPIClient.CreateClientResp{}, grpc.Errorf(400, "Secret Key invalid")
	}
	// 创建client
	clientUUID := uuid.NewV4().String()

	client := &DBModel.Client{Type: uint8(req.ClientType), UUID: clientUUID}
	result := MysqlDB.DB.Create(client)
	if result.Error != nil {
		return &pbBaseAPIClient.CreateClientResp{}, grpc.Errorf(400, "mysql insert error ")
	}
	// unit unit32 uni64 int 等类型后期需要统一处理一下
	return &pbBaseAPIClient.CreateClientResp{ClientID: client.ID}, nil
}

func (rpc *rpcBaseAPIClient) CreateClientToMessage(_ context.Context, req *pbBaseAPIClient.CreateClientToMessageReq) (*pbBaseAPIClient.CreateClientToMessaageResp, error) {
	client_to_msg := DBModel.NewClientToMessage()
	client_to_msg_id := client_to_msg.CreateClientToMessage(MysqlDB.DB, req.ClientID, req.RecvID)
	resp := &pbBaseAPIClient.CreateClientToMessaageResp{ClientToMsgID: client_to_msg_id}
	return resp, nil
}

func (rpc *rpcBaseAPIClient) CreateCollectToMessage(_ context.Context, req *pbBaseAPIClient.CreateCollectToMessageReq) (*pbBaseAPIClient.CreateCollectToMessageResp, error) {
	collect_to_msg := DBModel.NewCollectToMessage()
	collect_to_msg.CreateCollectToMessage(MysqlDB.DB)
	resp := &pbBaseAPIClient.CreateCollectToMessageResp{CollectToMsgID: 200}
	return resp, nil
}

func (rpc *rpcBaseAPIClient) Run() {

	address := "0.0.0.0" + ":" + strconv.Itoa(rpc.rpcPort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// pbMsgGateway.RegisterMsgGatewayServer(s, &rpcMsgGateway{})
	pbBaseAPIClient.RegisterBaseAPIClientServer(s, &rpcBaseAPIClient{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
