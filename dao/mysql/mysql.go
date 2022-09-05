package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() (err error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
