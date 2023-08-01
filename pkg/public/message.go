package public

// ======================WServer===============================
// websocket接收到的消息
type CommonMsgReq struct {
	MessageType uint `json:"message_type"`
}

type SingleMsgToWS struct {
}

type GroupMsgToWS struct {
}

type RelayMsgToWS struct {
}

type PullMsgToWS struct {
}

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
