package DBMysql

import (
	"fmt"


	gomysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}


type MysqlDB struct {
	//sync.RWMutex
	// 可能会有读写锁问题
	DB *gorm.DB
}

func (m *MysqlDB) InitMysqlDB() {
	dsn := "root:xinxin@5102G@tcp(home.xinxinblog.top:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gomysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error")
	}

	db.AutoMigrate(&User{})

	db.Set("gorm:table_options", "CHARSET=utf8")
	db.Set("gorm:table_options", "collation=utf8_unicode_ci")

	m.DB = db

	user := User{Name: "Tom"} 

	result := db.Create(&user) // pass pointer of data to Create
	fmt.Println(result)
}

// func (m *MysqlDB) DefaultGormDB() *gorm.DB {
// 	return m.db
// }
