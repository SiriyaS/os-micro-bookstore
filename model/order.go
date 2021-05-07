package model

import (
	"os-micro-bookstore/database"
	"os-micro-bookstore/form"
)

type OrderModel struct{}

func (om OrderModel) Add(order form.Order) (err error) {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return err
	}
	defer database.CloseConnection(conn)

	if err = conn.
		Table("order_main").
		Create(&order.Header).Error; err != nil {
		return
	}

	if err = conn.
		Table("order_detail").
		Create(&order.Details).Error; err != nil {
		return
	}

	return
}

func (om OrderModel) ReadByOrderNo(orderNo string) (form.Order, error) {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return form.Order{}, err
	}
	defer database.CloseConnection(conn)

	var order form.Order
	if err = conn.
		Table("order_main").
		Where("order_no = ?", orderNo).
		Find(&order.Header).Error; err != nil {
		return form.Order{}, err
	}

	if err = conn.
		Table("order_detail").
		Where("order_no = ?", orderNo).
		Find(&order.Details).Error; err != nil {
		return form.Order{}, err
	}

	return order, nil
}
