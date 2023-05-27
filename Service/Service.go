package Service

import (
	"230525/Struct"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

// InitSQLConn 初始化sql连接
func InitSQLConn() error {
	sqlDsn := "root:123456@tcp(127.0.0.1:3306)/student?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(sqlDsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&Struct.User{})
	if err != nil {
		return err
	}
	dbConn = db
	return nil
}
