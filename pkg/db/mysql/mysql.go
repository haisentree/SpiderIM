package DBMysql

import (
	"fmt"

	gomysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	//sync.RWMutex
	// 可能会有读写锁问题
	DB *gorm.DB
}

func (m *MysqlDB) InitMysqlDB() {
	dsn := "root:123@tcp(192.168.45.128:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gomysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error")
	}

	db.Set("gorm:table_options", "CHARSET=utf8")
	db.Set("gorm:table_options", "collation=utf8_unicode_ci")

	m.DB = db

}
