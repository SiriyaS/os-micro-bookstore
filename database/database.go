package database

import (
	"fmt"

	"os-micro-bookstore/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbAddress string

func Init() {
	conf := config.GetConfig()
	dbAddress = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		conf.GetString("database.user"),
		conf.GetString("database.password"),
		conf.GetString("database.host"),
		conf.GetString("database.port"),
		conf.GetString("database.initdb"))
}

func NewConnection() (*gorm.DB, error) {
	// connect to database
	conn, err := gorm.Open(mysql.Open(dbAddress), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func CloseConnection(conn *gorm.DB) error {
	// change conn to type *sql.DB to close the connection
	sqlDB, err := conn.DB()
	err = sqlDB.Close()
	return err
}
