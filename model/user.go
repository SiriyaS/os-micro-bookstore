package model

import (
	"fmt"
	"os-micro-bookstore/database"
)

type UserModel struct{}

func (um UserModel) Add(name string, email string, address string, tel string, username string, password string) error {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return err
	}
	defer database.CloseConnection(conn)

	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(address)
	fmt.Println(tel)

	return nil
}
