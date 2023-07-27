package DB

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	ID         uint64 `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	ClientUUID string         `gorm:unique`
	ClientType int32
}

func (Client) TableName() string {
	return "client"
}
