package DBModel

import (
	pkgError "SpiderIM/pkg/public/error"
	"log"

	"gorm.io/gorm"
)

// ==================================================ClientToMessage=====================================================================
// 存储client与client之间的消息关系
type ClientToMessage struct {
	Base
	ClientID       uint64
	RecvID         uint64
	MinSeq         uint64
	MaxSeq         uint64
	ClientMessages []ClientMessage
}

func (ClientToMessage) TableName() string {
	return "client_to_message"
}
func NewClientToMessage() *ClientToMessage {
	c := &ClientToMessage{}
	return c
}

// 创建ClientMessage
func (c *ClientToMessage) CreateClientToMessage(db *gorm.DB, client_id uint64, recv_id uint64) uint64 {
	c.ClientID = client_id
	c.RecvID = recv_id
	c.MinSeq = 1
	c.MaxSeq = 1
	result := db.Create(c)
	if result.Error != nil {
		log.Println(pkgError.Mysql_CreateClientMessage_Error)
	}
	return c.ID
}

func (c *ClientToMessage) FindByClientIDAndRecvID(db *gorm.DB, client_id uint64, recv_id uint64) ClientToMessage {
	var client_to_message ClientToMessage
	db.Where("client_id = ? AND recv_id = ?", client_id, recv_id).First(&client_to_message)
	return client_to_message
}

func (c *ClientToMessage) FindMaxSeqByID(db *gorm.DB, client_to_message_id uint64) uint64 {
	var client_to_message ClientToMessage
	result := db.First(&client_to_message, client_to_message_id)
	if result.Error != nil {
		log.Println("err 3234423")
	}
	return client_to_message.MaxSeq
}

func (c *ClientToMessage) IncMaxSeq(db *gorm.DB, client_to_message_id uint64) {
	var client_to_message ClientToMessage
	r1 := db.First(&client_to_message, client_to_message_id)
	if r1.Error != nil {
		log.Println("err 323asdds23")
	}
	client_to_message.MaxSeq = client_to_message.MaxSeq + 1
	r2 := db.Save(client_to_message)
	if r2.Error != nil {
		log.Println("err 32ddds23")
	}
}

// ===================================================ClientMessage=========================================================
type ClientMessage struct {
	Base
	ClientToMessageID uint64
	SeqID             uint64 `gorm:"index;unique"`
	Content           string
	IsSender          bool
}

func (ClientMessage) TableName() string {
	return "client_message"
}
func NewClientMessage() *ClientMessage {
	m := &ClientMessage{}
	return m
}

func (m *ClientMessage) CreateMessage(db *gorm.DB, client_message_id uint64, seq uint64, content string, is_sender bool) {
	m.ClientToMessageID = client_message_id
	m.SeqID = seq
	m.Content = content
	m.IsSender = is_sender
	result := db.Create(m)
	if result.Error != nil {
		log.Println("err 123")
	}
}

func (m *ClientMessage) FindMessageBySeq(db *gorm.DB, seq_start uint64, seq_end uint64) []ClientMessage {
	var client_messages []ClientMessage
	db.Where("seq_id >= ? AND seq_id <= ?", seq_start, seq_end).Find(&client_messages)
	return client_messages
}

// ==================================================CollectToMessage==========================================================
type CollectToMessage struct {
	Base
	MinSeq          uint64
	MaxSeq          uint64
	CollectMessages []CollectMessage
}

func (CollectToMessage) TableName() string {
	return "collect_to_message"
}

func NewCollectToMessage() *CollectToMessage {
	c := &CollectToMessage{}
	return c
}

func (c *CollectToMessage) CreateCollectToMessage(db *gorm.DB) {
	c.MinSeq = 1
	c.MaxSeq = 1
	result := db.Create(c)
	if result.Error != nil {
		log.Println("err 2132")
	}
}

func (c *CollectToMessage) FindByCollectID(db *gorm.DB, collect_to_msg_id uint64) CollectToMessage {
	var collect_to_message CollectToMessage
	result := db.First(&collect_to_message, collect_to_msg_id)
	if result.Error != nil {
		log.Println("23546")
	}
	return collect_to_message
}

func (c *CollectToMessage) IncMaxseq(db *gorm.DB, collect_to_msg_id uint64) {
	var collect_to_message CollectToMessage
	r1 := db.First(&collect_to_message, collect_to_msg_id)
	if r1.Error != nil {
		log.Println("23546")
	}
	collect_to_message.MaxSeq = collect_to_message.MaxSeq + 1
	r2 := db.Save(collect_to_message)
	if r2.Error != nil {
		log.Println("err 32ddds23")
	}
}

// ==================================================CollectMessage==========================================================
type CollectMessage struct {
	Base
	CollectToMessageID uint64
	SeqID              uint64 `gorm:"index;unique"`
	SendID             uint64
	Content            string
}

func (CollectMessage) TableName() string {
	return "collect_message"
}

func NewCollectMessage() *CollectMessage {
	c := &CollectMessage{}
	return c
}

func (c *CollectMessage) CreateCollectMessage(db *gorm.DB, collect_to_message_id uint64, content string, seq uint64, sender uint64) {
	c.CollectToMessageID = collect_to_message_id
	c.SeqID = seq
	c.SendID = sender
	c.Content = content
	result := db.Create(c)
	if result.Error != nil {
		log.Println("err 1123424")
	}
}

func (c *CollectMessage) FindMessageBySeq(db *gorm.DB, collect_to_message_id uint64, seq_start uint64, seq_end uint64) []CollectMessage {
	var collect_messages []CollectMessage
	db.Where("seq_id >= ? AND seq_id <= ?", seq_start, seq_end).Find(&collect_messages)
	return collect_messages
}
