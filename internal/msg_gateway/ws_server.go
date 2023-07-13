package gate

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"SpiderIM/pkg/utils"
)

type WSClient struct {
	*websocket.Conn
	PlatformID int32
	w          *sync.Mutex
	// token        string			// 用户和使用clinet的用户信息关联
	clientID string
}

type WServer struct {
	wsAddr       string
	wsMaxConnNum int
	wsUpGrader   *websocket.Upgrader
	wsUserToConn map[string]map[int]*WSClient
}

func (ws *WServer) OnInit(wsPort int) {
	ws.wsAddr = ":" + utils.IntToString(wsPort)
	ws.wsMaxConnNum = 20
	ws.wsUserToConn = make(map[string]map[int]*WSClient)
	ws.wsUpGrader = &websocket.Upgrader{
		HandshakeTimeout: 5 * time.Second,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		CheckOrigin:      func(r *http.Request) bool { return true },
	}
}

func (ws *WServer) Run() {
	http.HandleFunc("/", ws.wsHandler)
	err := http.ListenAndServe(ws.wsAddr, nil)
	if err != nil {
		panic("Ws listening err:" + err.Error())
	}
}

func (ws *WServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.wsUpGrader.Upgrade(w, r, nil)
	defer conn.Close()
	if err != nil {
		log.Println(err)
		return
	}
	for {
		//messageType int, p []byte, err error
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func (ws *WServer) writeMsg(conn *WSClient, msgType int, msg []byte) error {
	conn.w.Lock()
	defer conn.w.Unlock()
	// conn.SetWriteDeadline(time.Now().Add(time.Duration(60) * time.Second))
	return conn.WriteMessage(msgType, msg)
}

func (ws *WServer) readMsg() {

}

func (ws *WServer) addClientConn() {

}

func (ws *WSClient) delClientConn() {

}

//======================功能函数=============================
