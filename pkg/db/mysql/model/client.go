package DBModel

import (
	"log"

	pkgError "SpiderIM/pkg/public/error"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// ============================================================================================================
type Client struct {
	Base
	UUID           string `gorm:"unique"`
	Type           uint8
	ClientToMessages []ClientToMessage
}

func (Client) TableName() string {
	return "client"
}

func NewClient() *Client {
	c := &Client{}
	return c
}

// 创建Client
func (c *Client) CreateClient(db *gorm.DB, client_type uint8) {
	c.UUID = uuid.NewV4().String()
	c.Type = client_type
	result := db.Create(c)
	if result.Error != nil {
		log.Println(pkgError.Mysql_CreateClient_Error)
	}
}
