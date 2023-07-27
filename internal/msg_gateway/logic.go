package MsgGateway

import (
	pbMsgGateway "SpiderIM/pkg/proto/msg_gateway"
	"context"
	"encoding/json"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
)

/*消息说明
{
	"sendID":12,
	"clientType":10	// clientMsg
	"platformID":2,
	"recvID":2,
	"content":
}
push模块推送消息，直接发送，不经过处理 // serverMsg
{
	"sendID":12,
	"clientType":1,
	"platformID":2,
	"recvID":2,
	"content": "dsafasdfa"
}
*/

// 消息中不需要携带SendID等信息，conn中会包含
type CommonMsg struct {
	// SendID     int    `json:"sendID" validate:"required"`
	MessageType int    `json:"MessageType" validate:"required"`
	PlatformID  int    `json:"platformID" validate:"required"`
	ReceID      int    `json:"recvID" validate:"required"`
	Content     string `json:"content" validate:"required"`
}

func (ws *WServer) msgParse(conn *WSClient, binaryMsg []byte) {
	msg := CommonMsg{}
	json.Unmarshal(binaryMsg, &msg)
	validate := validator.New()
	if err := validate.Struct(msg); err != nil {
		log.Println("validate error:", err)
		return
	}
	switch msg.MessageType {
	// 此消息直接发送
	case 1:
		log.Println("serverMsg")
		// 此处之前出错了，
		err := ws.writeMsg(conn, websocket.TextMessage, binaryMsg)
		// err := ws.writeMsg(conn, websocket.BinaryMessage, binaryMsg)
		if err != nil {
			log.Println("case1 error:", err)
		}
		log.Println("writeMsg SendID:", conn.clientID)
	// 此消息发送给后端
	case 10:
		// 	req := &pbMsgGateway.MessageReq{
		// 		Type:    "1",
		// 		Message: string(message),
		// 	}

		// 	// 将收到的消息发送给rpc后端
		// 	resp, err := MsgGatewaySrvClient.ReceiveMessage(context.Background(), req)
		// 	if err != nil {
		// 		log.Print("test message fail")
		// 	}
		// 	// temp := resp.ErrMsg + "test"
		// 	// message2 := []byte(temp)
		// 	log.Print(resp.ErrMsg)
		req := &pbMsgGateway.ClientMsgReq{
			SendID:  int64(conn.clientID), // conn
			ReceID:  int64(msg.ReceID),
			Content: msg.Content,
		}
		resp, err := MsgGatewaySrvClient.ReceiveClientMsg(context.Background(), req)
		if err != nil {
			log.Println("case 10 error")
		}
		log.Println("case 10 resp:", resp.ErrMsg)
		log.Println("clientMsg")
	default:
		log.Println("clientType error")
	}
}
