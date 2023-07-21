package MsgGateway

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"SpiderIM/pkg/utils"
)

// 客户端类型说明
// 1: 内部push客户端
// 10: 普通用户客户端
// 20: 群组客户端

// 平台说明
// 0:内部服务
// 1:手机
// 2:PC
// 3:web
type WSClient struct {
	*websocket.Conn
	platformID int32
	w          *sync.Mutex
	// token        string			// 用户和使用clinet的用户信息关联
	clientID   string
	clinetType int32
}

type WServer struct {
	wsAddr         string
	wsMaxConnNum   int
	wsUpGrader     *websocket.Upgrader
	wsClientToConn map[string]map[int32]*WSClient // map[client_id][platform_id]
}

func (ws *WServer) OnInit(wsPort int) {
	ws.wsAddr = ":" + utils.IntToString(wsPort)
	ws.wsMaxConnNum = 20
	ws.wsClientToConn = make(map[string]map[int32]*WSClient)
	ws.wsUpGrader = &websocket.Upgrader{
		HandshakeTimeout: 5 * time.Second,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		CheckOrigin:      func(r *http.Request) bool { return true },
	}
}

func (ws *WServer) Run() {
	http.HandleFunc("/ws", ws.wsHandler)
	err := http.ListenAndServe(ws.wsAddr, nil)
	if err != nil {
		panic("Ws listening err:" + err.Error())
	}
}

// 进行websocket连接，连接的时候会进行信息的认证
func (ws *WServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.wsUpGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	newConn := &WSClient{Conn: conn, platformID: 1, clientID: "121", clinetType: 10}
	ws.addClientConn(newConn)
	// 用户连接成功，对连接的数据进行读取
	go ws.readMsg(newConn)

	// for {
	// 	//messageType int, p []byte, err error
	// 	mt, message, err := conn.ReadMessage()
	// 	if err != nil {
	// 		log.Println("read:", err)
	// 		break
	// 	}
	// 	log.Printf("recv: %s", message)

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
	// err = conn.WriteMessage(websocket.BinaryMessage, []byte("test msg"))
	// if err != nil {
	// 	log.Println("write:", err)
	// }
	// }
}

func (ws *WServer) writeMsg(conn *WSClient, msgType int, message []byte) error {
	// conn.w.Lock()
	// defer conn.w.Unlock()
	conn.SetWriteDeadline(time.Now().Add(time.Duration(60) * time.Second))
	return conn.WriteMessage(msgType, message)
}

func (ws *WServer) readMsg(conn *WSClient) {
	for {
		//messageType int, p []byte, err error
		msgType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		// 对接收的消息进行处理
		log.Printf("recv: %s", message)
		log.Printf("msgType: %d", msgType)
		err = conn.WriteMessage(websocket.TextMessage, []byte("send"))
		if err != nil {
			log.Println("readMsg send error:", err)
		}
		ws.msgParse(conn, message)

		// req := &pbMsgGateway.MessageReq{
		// 	Type:    "1",
		// 	Message: string(message),
		// }
	}
}

func (ws *WServer) addClientConn(conn *WSClient) {
	i := make(map[int32]*WSClient)
	i[conn.platformID] = conn
	ws.wsClientToConn[conn.clientID] = i
}

func (ws *WSClient) delClientConn(conn *WSClient) {
	conn.Close()
}

//======================功能函数=============================
