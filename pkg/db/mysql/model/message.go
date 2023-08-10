package DBModel

import (
	pkgError "SpiderIM/pkg/public/error"
	"log"

	"gorm.io/gorm"
)

// =======================================================================================================================
// 存储client与client之间的消息关系
type ClientToMessage struct {
	Base
	ClientID uint64
	RecvID   uint64
	MinSeq   uint64
	MaxSeq   uint64
}

func (ClientToMessage) TableName() string {
	return "client_to_message"
}
func NewClientToMessage() *ClientToMessage {
	c := &ClientToMessage{}
	return c
}

// 创建ClientMessage
func (c *ClientToMessage) CreateClientToMessage(db *gorm.DB, client_id uint64, recv_id uint64) {
	c.ClientID = client_id
	c.RecvID = recv_id
	c.MinSeq = 1
	c.MaxSeq = 1
	result := db.Create(c)
	if result.Error != nil {
		log.Println(pkgError.Mysql_CreateClientMessage_Error)
	}
}

func (c *ClientToMessage) FindByClientIDAndRecvID(db *gorm.DB, client_id uint64, recv_id uint64) uint64 {
	var client_to_message ClientToMessage
	db.Where("client_id = ? AND recv_id = ?", client_id, recv_id).First(&client_to_message)
	return client_to_message.ID
}

// ============================================================================================================
type ClientMessage struct {
	Base
	ClientToMessageID uint64
	SeqID             uint64 `gorm:"index;unique"`
	Content           string
}

func (ClientMessage) TableName() string {
	return "client_message"
}
func NewClientMessage() *ClientMessage {
	m := &ClientMessage{}
	return m
}

func (m *ClientMessage) CreateMessage(db *gorm.DB, client_message_id uint64, seq uint64, content string) {
	m.ClientToMessageID = client_message_id
	m.SeqID = seq
	m.Content = content
	result := db.Create(m)
	if result.Error != nil {
		log.Println("err 123")
	}
}

// ============================================================================================================
type GroupToMessage struct {
	Base
}
