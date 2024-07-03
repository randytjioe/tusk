package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 3310
	user     = "root"
	password = ""
	dbName   = "tusk"
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbName,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return database
}
