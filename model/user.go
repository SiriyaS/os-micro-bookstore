package model

import (
	"os-micro-bookstore/database"
	"os-micro-bookstore/form"
)

type UserModel struct{}

func (um UserModel) Add(user form.UserInfoRequest) (err error) {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return err
	}
	defer database.CloseConnection(conn)

	if err = conn.
		Table("users").
		Create(&user).Error; err != nil {
		return
	}

	return
}

func (um UserModel) ReadByEmail(email string) (user form.UserInfo, err error) {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return form.UserInfo{}, err
	}
	defer database.CloseConnection(conn)

	if err = conn.
		Table("users").
		Where("email = ?", email).
		Find(&user).Error; err != nil {
		return
	}

	return
}

func (um UserModel) UpdateByEmail(email string, user form.UserInfoRequest) (err error) {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return err
	}
	defer database.CloseConnection(conn)

	if user.Name != "" {
		if err = conn.
			Table("users").
			Where("email = ?", email).
			Update("name", user.Name).Error; err != nil {
			return
		}
	}
	if user.Address != "" {
		if err = conn.
			Table("users").
			Where("email = ?", email).
			Update("address", user.Address).Error; err != nil {
			return
		}
	}
	if user.Telephone != "" {
		if err = conn.
			Table("users").
			Where("email = ?", email).
			Update("telephone", user.Telephone).Error; err != nil {
			return
		}
	}
	if user.Username != "" {
		if err = conn.
			Table("users").
			Where("email = ?", email).
			Update("username", user.Username).Error; err != nil {
			return
		}
	}

	return nil
}
