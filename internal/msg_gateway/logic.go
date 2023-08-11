package MsgGateway

import (
	pbMsgGateway "SpiderIM/pkg/proto/msg_gateway"
	pkgMessage "SpiderIM/pkg/public/message"
	"context"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

func (ws *WServer) msgParse(conn *WSClient, binaryMsg []byte) {
	m := pkgMessage.CommonMsg{}
	json.Unmarshal(binaryMsg, &m)

	if err := Validate.Struct(m); err != nil {
		log.Println("validate error:", err)
		return
	}
	switch m.MessageType {
	case pkgMessage.Single_Common_Message_Request:
		log.Println("msg prase single common message")
		ws.parseSingleCommMsg(conn, &m)
	case pkgMessage.Single_Relay_Message_Request:
		log.Println("serverMsg")
		ws.parseSingleRelayMsg(conn, &m)
	case pkgMessage.Group_Common_Message_Request:
		log.Println("group msg")
		ws.paraGroupCommMsg(conn,&m)
	default:
		log.Println("clientType error")
	}
}

func (ws *WServer) parseSingleCommMsg(conn *WSClient, msg *pkgMessage.CommonMsg) {
	d := &pkgMessage.SingleCommMsgReq{}
	json.Unmarshal([]byte(msg.Data), &d)
	if err := Validate.Struct(d); err != nil {
		log.Println("validate error: 1423", err)
		return
	}

	req := &pbMsgGateway.SingleMsgReq{
		SendID:  msg.SendID,
		RecvID:  d.RecvID,
		Content: d.Content,
	}

	resp, err := MsgGatewaySrvClient.ReceiveSingleMsg(context.Background(), req)
	if err != nil {
		log.Println("case 10afs error")
	}

	log.Println("clientMsg case 10 resp:", resp.Message)
}

func (ws *WServer) parseSingleRelayMsg(conn *WSClient, msg *pkgMessage.CommonMsg) {
	d := &pkgMessage.SingleRelayMsgReq{}
	json.Unmarshal([]byte(msg.Data), &d)
	if err := Validate.Struct(d); err != nil {
		log.Println("validate error: 14dd23", err)
		return
	}

	if err := ws.writeMsg(conn, websocket.TextMessage, []byte(d.Content)); err != nil {
		log.Println("case1 error:", err)
	}
}

func (ws *WServer) paraGroupCommMsg(conn *WSClient,msg *pkgMessage.CommonMsg) {

}
