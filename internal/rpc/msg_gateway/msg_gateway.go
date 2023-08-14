package rpcMsgGateway

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"

	DBMysql "SpiderIM/pkg/db/mysql"
	DBModel "SpiderIM/pkg/db/mysql/model"
	pbMsgGateway "SpiderIM/pkg/proto/msg_gateway"
	pkgMessage "SpiderIM/pkg/public/message"
	"SpiderIM/pkg/rabbitmq"
)

var (
	MessageProducer *rabbitmq.Producer
	MysqlDB         DBMysql.MysqlDB
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
	MysqlDB.InitMysqlDB()
	MysqlDB.DB.AutoMigrate(&DBModel.Client{}, &DBModel.ClientToMessage{}, &DBModel.ClientMessage{})
}

func (rpc *rpcMsgGateway) ReceiveSingleMsg(_ context.Context, req *pbMsgGateway.SingleMsgReq) (*pbMsgGateway.SingleMsgResp, error) {
	log.Println("receiveClientMsg:", req.Content)
	// 1.给消息分配seq
	var seq uint64
	switch req.MsgType {
	case pkgMessage.Single_Common_Message_Request:
		client_to_msg := DBModel.NewClientToMessage()
		temp := client_to_msg.FindByClientIDAndRecvID(MysqlDB.DB, req.SendID, req.RecvID)
		seq = temp.MaxSeq
		client_to_msg.IncMaxSeq(MysqlDB.DB, temp.ID)

		// 2.将消息存储到数据库中,一条client消息sender与recver都要存储一遍
		client_msg := DBModel.NewClientMessage()
		client_msg.CreateMessage(MysqlDB.DB, temp.ID, seq, req.Content, true)
		temp = client_to_msg.FindByClientIDAndRecvID(MysqlDB.DB, req.RecvID, req.SendID)
		client_msg.CreateMessage(MysqlDB.DB, temp.ID, seq, req.Content, false)

	case pkgMessage.Group_Common_Message_Request:
		collect_to_msg := DBModel.NewCollectToMessage()
		temp := collect_to_msg.FindByCollectID(MysqlDB.DB, req.RecvID)
		seq = temp.MaxSeq
		collect_to_msg.IncMaxseq(MysqlDB.DB, temp.ID)

		// 2.将消息存储到数据库中,一条group消息只需要存储一遍
		collect_msg := DBModel.NewCollectMessage()
		collect_msg.CreateCollectMessage(MysqlDB.DB, temp.ID, req.Content, seq, req.SendID)
	default:
		seq = 0
	}

	// 2.将消息存入MQ中
	m := &pkgMessage.SingleMsgToMQ{
		SendID:  req.SendID,
		RecvID:  req.RecvID,
		SeqID:   seq,
		MsgType: req.MsgType,
		Content: req.Content,

		CreateTime: time.Now().Unix(), // 与存储的create时间不完全相同
	}

	j, err := json.Marshal(&m)
	if err != nil {
		log.Println("json marshal fail")
		return nil, grpc.Errorf(500, "json marshal fail")
	}

	MessageProducer.SendMessage(j)

	return &pbMsgGateway.SingleMsgResp{Code: 200, Message: "success"}, nil
}

func (rpc *rpcMsgGateway) ReceiveListMsg(_ context.Context, req *pbMsgGateway.ListMsgReq) (*pbMsgGateway.ListMsgResp, error) {
	switch req.MsgType {
	case pkgMessage.Group_List_Message_Request:
		// 遍历RecvList直接,不存储到数据库，直接发送给MQ进行转发
		for _, v := range req.RecvID {
			m := &pkgMessage.SingleMsgToMQ{
				SendID:  req.SendID,
				RecvID:  v,
				SeqID:   req.SeqID,
				MsgType: req.MsgType,
				Content: req.Content,

				CreateTime: time.Now().Unix(), // 与存储的create时间不完全相同
			}

			j, err := json.Marshal(&m)
			if err != nil {
				log.Println("json marshal fail")
				return nil, grpc.Errorf(500, "json marshal fail")
			}

			MessageProducer.SendMessage(j)
		}

	default:
		log.Println("msg Type Error")
	}
	return &pbMsgGateway.ListMsgResp{Code: 200, Message: "success"}, nil
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
