package public

// ======================宏定义================================
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
	Single_Common_Message_Request = 1
	Group_Common_Message_Request  = 2
)

// ======================WServer===============================
// websocket接收到的请求消息
type CommonMsg struct {
	MessageType uint8  `json:"message_type"`
	PlatformID  uint8  `json:"platform_id"`
	Data        string `json:"data"`
}

type SingleMsgReq struct {
}

type GroupMsgReq struct {
}

type RelayMsgReq struct {
}

type PullMsgReq struct {
}

// websocket接收到的响应消息

// =====================MQ=======================================
// websocket存入MQ中的消息

type MessageToMQ struct {
	SendID     int64  `json:"send_id"`
	RecvID     int64  `json:"recv_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

// type MQSendMessage struct {
// 	SendID     int64  `json:"send_id"`
// 	RecvID     int64  `json:"recv_id"`
// 	Content    string `json:"content"`
// 	CreateTime int64  `json:"create_time"`
// }
