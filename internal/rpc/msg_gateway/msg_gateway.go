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
	DBRedis "SpiderIM/pkg/db/redis"
	pbMsgGateway "SpiderIM/pkg/proto/msg_gateway"
	pkgMessage "SpiderIM/pkg/public/message"
	"SpiderIM/pkg/rabbitmq"
)

var (
	MessageProducer *rabbitmq.Producer
	MysqlDB         DBMysql.MysqlDB
	RedisDB         DBRedis.RedisDB
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
	MysqlDB.DB.AutoMigrate(&DBModel.Client{}, &DBModel.ClientToMessage{}, &DBModel.ClientMessage{},
		&DBModel.CollectToMessage{}, &DBModel.CollectMessage{})
	RedisDB.InitRedisDB()
}

func (rpc *rpcMsgGateway) ReceiveSingleMsg(_ context.Context, req *pbMsgGateway.SingleMsgReq) (*pbMsgGateway.SingleMsgResp, error) {
	log.Println("receiveClientMsg:", req.Content)
	// 1.给消息分配seq
	var seq uint64
	switch req.MsgType {
	// 该功能过程较多，需要讨论失败带来的错误情况，以及以后
	case pkgMessage.Single_Common_Message_Request:
		client_to_msg := DBModel.NewClientToMessage()
		temp := client_to_msg.FindByClientIDAndRecvID(MysqlDB.DB, req.SendID, req.RecvID)
		seq = temp.MaxSeq
		client_to_msg.IncMaxSeq(MysqlDB.DB, temp.ID)

		// 2.将消息存储到数据库中,一条client消息sender与recver都要存储一遍
		client_msg := DBModel.NewClientMessage()
		client_msg.CreateMessage(MysqlDB.DB, temp.ID, seq, req.Content, true)
		temp = client_to_msg.FindByClientIDAndRecvID(MysqlDB.DB, req.RecvID, req.SendID)
		log.Println("temp:", temp)
		client_msg2 := DBModel.NewClientMessage()
		client_msg2.CreateMessage(MysqlDB.DB, temp.ID, seq, req.Content, false)
		client_to_msg2 := DBModel.NewClientToMessage()
		client_to_msg2.IncMaxSeq(MysqlDB.DB, temp.ID)

	case pkgMessage.Group_Common_Message_Request:
		collect_to_msg := DBModel.NewCollectToMessage()
		temp := collect_to_msg.FindByCollectID(MysqlDB.DB, req.RecvID)
		seq = temp.MaxSeq
		collect_to_msg.IncMaxseq(MysqlDB.DB, temp.ID)

		// 2.将消息存储到数据库中,一条group消息只需要存储一遍
		collect_msg := DBModel.NewCollectMessage()
		collect_msg.CreateCollectMessage(MysqlDB.DB, temp.ID, req.Content, seq, req.SendID)

		return &pbMsgGateway.SingleMsgResp{Code: 200, Message: "success"}, nil
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
			log.Println("List Message Divide into MQ")

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

func (rpc *rpcMsgGateway) ControlPullClientMsg(_ context.Context, req *pbMsgGateway.PullClientMsgReq) (*pbMsgGateway.PullClientMsgResp, error) {
	client_to_msg := DBModel.NewClientToMessage()
	client_msg := DBModel.NewClientMessage()
	resp := &pbMsgGateway.PullClientMsgResp{
		Code: 200,
	}
	for _, v := range req.ClientToSeq {
		find_client_to_msg := client_to_msg.FindByClientIDAndRecvID(MysqlDB.DB, req.OwnerID, v.ClientID)
		if v.SeqID < find_client_to_msg.MaxSeq {
			m := client_msg.FindMessageBySeq(MysqlDB.DB, find_client_to_msg.ID, v.SeqID, find_client_to_msg.MaxSeq)
			for _, val := range m {
				temp := pbMsgGateway.CommonClientToMsg{
					SeqID:      val.SeqID,
					OwnerID:    req.OwnerID,
					ClientID:   v.ClientID,
					IsSneder:   val.IsSender,
					CreateTime: val.CreatedAt.Unix(),
					Content:    val.Content,
				}
				resp.ClientToMsg = append(resp.ClientToMsg, &temp)
			}
		}
	}
	return resp, nil
}

func (rpc *rpcMsgGateway) ControlPullCollectMsg(_ context.Context, req *pbMsgGateway.PullCollectMsgReq) (*pbMsgGateway.PullCollectMsgResp, error) {
	collect_to_msg := DBModel.NewCollectToMessage()
	collect_msg := DBModel.NewCollectMessage()
	resp := &pbMsgGateway.PullCollectMsgResp{
		Code: 200,
	}
	for _, v := range req.CollectToSeq {
		log.Println("CollectID:", v.CollectID)
		find_collect_to_msg := collect_to_msg.FindByCollectID(MysqlDB.DB, v.CollectID)
		if v.SeqID < find_collect_to_msg.MaxSeq {
			m := collect_msg.FindMessageBySeq(MysqlDB.DB, v.CollectID, v.SeqID, find_collect_to_msg.MaxSeq)
			for _, val := range m {
				temp := pbMsgGateway.CommonCollectToMsg{
					SeqID:      val.SeqID,
					CollectID:  v.CollectID,
					SendID:     val.SendID,
					CreateTime: val.CreatedAt.Unix(),
					Content:    val.Content,
				}
				resp.CollectToMsg = append(resp.CollectToMsg, &temp)
			}
		}
	}

	return resp, nil
}

func (rpc *rpcMsgGateway) ControlGetClientMaxSeq(_ context.Context, req *pbMsgGateway.GetClientMaxSeqReq) (*pbMsgGateway.GetClientMaxSeqResp, error) {
	client_to_msg := DBModel.NewClientToMessage()
	resp := &pbMsgGateway.GetClientMaxSeqResp{}
	for _, v := range req.ClientList {
		find_client_to_msg := client_to_msg.FindByClientIDAndRecvID(MysqlDB.DB, req.OwnerID, v)
		temp := &pbMsgGateway.CommonClientToSeq{
			ClientID: v,
			SeqID:    find_client_to_msg.MaxSeq,
		}
		resp.ClientToSeq = append(resp.ClientToSeq, temp)
	}
	return resp, nil
}

func (rpc *rpcMsgGateway) ControlGetCollectMaxSeq(_ context.Context, req *pbMsgGateway.GetCollectMaxSeqReq) (*pbMsgGateway.GetCollectMaxSeqResp, error) {
	collect_to_msg := DBModel.NewCollectToMessage()
	resp := &pbMsgGateway.GetCollectMaxSeqResp{}
	for _, v := range req.CollectList {
		find_collect_to_msg := collect_to_msg.FindByCollectID(MysqlDB.DB, v)
		temp := &pbMsgGateway.CommonCollectToSeq{
			CollectID: v,
			SeqID:     find_collect_to_msg.MaxSeq,
		}
		resp.CollectToSeq = append(resp.CollectToSeq, temp)
	}

	return resp, nil
}

func (rpc *rpcMsgGateway) ControlGetClientStatus(_ context.Context, req *pbMsgGateway.GetClientStatusReq) (*pbMsgGateway.GetClientStatusResp, error) {

	resp := &pbMsgGateway.GetClientStatusResp{}

	for _, v := range req.ClientIDList {
		s := RedisDB.GetClientStauts(v)
		temp := &pbMsgGateway.CommonClientToStatus{
			ClientID: v,
			IsOnline: s,
		}
		resp.StatusList = append(resp.StatusList, temp)
	}
	return resp, nil
}

// 获取MaxSeq，用户对比本地数据库，按需拉去消息

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
