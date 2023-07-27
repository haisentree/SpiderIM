package public

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
