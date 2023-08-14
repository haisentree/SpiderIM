package pkgMessage

// ===============================================宏定义==============================================
// 平台类型
const (
	Internal_Platform = 0
	Android_Platform  = 1
	IOS_Platform      = 2
	PC_Platform       = 3
	Web_Platform      = 4
)

// 消息类型
const (
	Single_Common_Message_Request = 1 // client对client发送的单条消息
	Single_Relay_Message_Request  = 2 // msg_relay从MQ中读取，转发给client
	Group_Common_Message_Request  = 3 // 该条消息用于存储，不需要转发
	Group_List_Message_Request    = 4 // 该条消息包含recv数组，需要进行转发消息,自带seq
	Control_Pull_List_Message     = 5 // 首次登录，刷新时候，同步与数据库中的消息
	Control_Get_Max_Seq           = 6 // 获取seq，与本都seq对比
	Control_Get_Status            = 7 // 获取client的在线状态
)

// =============================================WServer===============================================
// websocket接收到的请求消息
type CommonMsg struct {
	MessageType uint8  `json:"message_type" validate:"required"`
	SendID      uint64 `json:"send_id" validate:"required"`
	PlatformID  uint8  `json:"platform_id" validate:"required"`
	Data        string `json:"data" validate:"required"`
}

type SingleCommMsgReq struct {
	RecvID  uint64 `json:"recv_id" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type SingleRelayMsgReq struct {
	RecvID  uint64 `json:"recv_id" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type GroupCommMsgReq struct {
	RecvID  uint64 `json:"recv_id" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type GroupListMsgReq struct {
	RecvIDList []uint64 `json:"recv_id" validate:"required"`
	SeqID      uint64   `json:"seq_id" validate:"required"`
	Content    string   `json:"content" validate:"required"`
}

// websocket接收到的响应消息

// =================================================MQ========================================================
// websocket存入MQ中的消息

type SingleMsgToMQ struct {
	SendID     uint64 `json:"send_id"`
	RecvID     uint64 `json:"recv_id"`
	SeqID      uint64 `json:"seq_id"`
	MsgType    uint32 `json:"msg_type"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
