package MsgGateway

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"

	DBModel "SpiderIM/pkg/db/mysql/model"
	pkgPublic "SpiderIM/pkg/public/message"
	"SpiderIM/pkg/utils"
)

// map存在并发安全，读去某个值时候，另一个一个goroutine更新该值，就会panic error
type WServer struct {
	wsAddr         string
	wsMaxConnNum   int
	wsUpGrader     *websocket.Upgrader
	wsClientToConn map[uint64]map[uint8]*WSClient
}

type WSClient struct {
	*websocket.Conn
	platformID uint8
	clientID   uint64
	clinetType uint8
}

func (ws *WServer) OnInit(wsPort int) {
	ws.wsAddr = ":" + utils.IntToString(wsPort)
	ws.wsMaxConnNum = 20
	ws.wsClientToConn = make(map[uint64]map[int32]*WSClient)
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

	// 1.提取clientID、clientUUID参数
	r.ParseForm()
	log.Println("request form:", r.Form)

	clientID, ok := r.Form["clientID"]
	if !ok {
		//此为不存在
		log.Println("clientID is none!")
		return
	}

	clientID_uint64, err := strconv.ParseUint(clientID[0], 10, 64)
	if err != nil {
		log.Println("clientID conv to int fail!")
		return
	}
	clientUUID, ok := r.Form["clientUUID"]
	if !ok {
		//此为不存在
		log.Println("clientUUID is none!")
		return
	}

	platformID, ok := r.Form["platformID"]
	if !ok {
		//此为不存在
		log.Println("platformID is none!")
		return
	}
	platformID_uint8, err := strconv.ParseUint(platformID[0], 10, 64)
	if err != nil {
		log.Println("clientID conv to int fail!")
		return
	}

	// 2.查询clientID、clientUUID
	dbClient := DBModel.NewClient()
	dbClient.FindByClientID(MysqlDB.DB, clientID_uint64)
	if dbClient.UUID != clientUUID[0] {
		log.Println("uuid error")
		return
	}

	// 3.校验plantformID是否合法
	switch platformID_uint8 {
	case 0, 1, 2, 3:
		log.Println("platform true")
	default:
		log.Println("platform error")
		return
	}
	// 4.建立websocket连接
	conn, err := ws.wsUpGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	newConn := &WSClient{Conn: conn, platformID: uint8(platformID_uint8), clientID: dbClient.ID, clinetType: dbClient.Type}
	ws.addClientConn(newConn)
	// 用户连接成功，对连接的数据进行读取
	go ws.readMsg(newConn)
}

func (ws *WServer) writeMsg(conn *WSClient, msgType int, message []byte) error {
	conn.SetWriteDeadline(time.Now().Add(time.Duration(60) * time.Second))
	return conn.WriteMessage(msgType, message)
}

func (ws *WServer) readMsg(conn *WSClient) {
	for {
		msgType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				// 连接正常关闭或正在关闭
				log.Println("连接关闭:", err)
				ws.delClientConn(conn)
			} else {
				// 连接异常关闭
				log.Println("连接异常关闭:", err)
				ws.delClientConn(conn)
			}
			log.Println("ws conn error", err)
			break
		}
		// 对接收的消息进行处理
		log.Printf("recv: %s", message)
		log.Printf("msgType: %d", msgType)
		log.Printf("platformID: %d", conn.platformID)
		err = conn.WriteMessage(websocket.TextMessage, []byte("send"))
		if err != nil {
			log.Println("readMsg send error:", err)
		}
		// 如果解析到platformID=0,直接发送给RecvID

		if conn.platformID == 0 {
			// 1.解析消息
			log.Println("recv platformID==1 message")
			var messageToMQ pkgPublic.MessageToMQ
			json.Unmarshal(message, &messageToMQ)
			// 2.查看RecvID是否连接在当前wss
			_, ok := ws.wsClientToConn[uint64(messageToMQ.RecvID)]
			if ok {
				log.Println("RecvID is online")
				// 3.发送消息
				for k, v := range ws.wsClientToConn[uint64(messageToMQ.RecvID)] {
					err := ws.writeMsg(v, websocket.TextMessage, message)
					log.Println("RecvID is sending platform:", v)
					if err != nil {
						log.Println("Sned RecvID error,platform:", k)
					}
				}

			} else {
				// 用户连接再其他的wss,通过redis判断用户是否在线
				// 后期再改
				log.Println("RecvID offonline")
			}
			// RecvID不在当前wss，message就丢弃

		} else {
			ws.msgParse(conn, message)
		}

		// req := &pbMsgGateway.MessageReq{
		// 	Type:    "1",
		// 	Message: string(message),
		// }
	}
}

func (ws *WServer) addClientConn(conn *WSClient) {
	i := make(map[uint8]*WSClient)
	i[conn.platformID] = conn
	ws.wsClientToConn[conn.clientID] = i
	RedisDB.SetClientStatus(conn.clientID, true)
	// reids 保存client在线状态
}

func (ws *WServer) delClientConn(conn *WSClient) {
	conn.Conn.Close() // 不知道conn已经关闭，执行这条语句会抛出错误吗？会抛出错误，但可以不接受
	delete(ws.wsClientToConn[conn.clientID], conn.platformID)
	if len(ws.wsClientToConn[conn.clientID]) == 0 {
		delete(ws.wsClientToConn, conn.clientID)
	}
	RedisDB.SetClientStatus(conn.clientID, false)
}
