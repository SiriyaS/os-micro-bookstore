package controller

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"

	"os-micro-bookstore/form"
	"os-micro-bookstore/model"
)

type OrderController struct{}

func (oc OrderController) CreateOrder(c *gin.Context) {
	log.Println("[Order: CreateOrder]")

	orderModel := model.OrderModel{}

	var request form.Order
	err := c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while binding to model.",
		})
		return
	}

	var orderNo string
	// generating orderNo
	for {
		// generate 5 digit number
		var num = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
		buf := make([]byte, 5)
		x, err := io.ReadAtLeast(rand.Reader, buf, 5)
		if x != 5 {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error while creating order",
			})
			return
		}
		for i := 0; i < len(buf); i++ {
			buf[i] = num[int(buf[i])%len(num)]
		}

		orderNo = fmt.Sprintf("ORDER%s", string(buf))

		// check if already exists
		emptyOrder := form.Order{
			Header:  form.OrderHeader{},
			Details: []form.OrderDetail{},
		}
		orderFound, err := orderModel.ReadByOrderNo(orderNo)
		// fmt.Printf("found\nvalue=%#v\n", orderFound)
		// fmt.Printf("struct\nvalue=%#v\n", emptyOrder)
		// fmt.Println(reflect.DeepEqual(orderFound, emptyOrder))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error while creating order",
			})
			return
		}
		if reflect.DeepEqual(orderFound, emptyOrder) {
			log.Println("orderNo has't existed. Create orderNo succcessfully.")
			break
		}
		log.Println("orderNo already exists. Generating new one...")
	}

	// calculate total and grandtotal
	var grandTotal float64
	for i := 0; i < len(request.Details); i++ {
		request.Details[i].Total = float64(request.Details[i].Quantity) * request.Details[i].UnitPrice
		grandTotal += request.Details[i].Total
		request.Details[i].OrderNo = orderNo
	}

	request.Header.OrderNo = orderNo
	request.Header.OrderDate = time.Now()
	request.Header.GrandTotal = grandTotal

	// fmt.Printf("%+v\n", request)

	err = orderModel.Add(request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while creating order",
		})
		return
	}

	log.Printf("Create order[%s] successfully", request.Header.OrderNo)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Create order[%s] successfully", request.Header.OrderNo),
	})
}

func (oc OrderController) GetOrderByOrderNo(c *gin.Context) {
	log.Println("[Book: GetOrderByOrderNo]")
	orderModel := model.OrderModel{}

	orderNo := c.Query("order_no")

	order, err := orderModel.ReadByOrderNo(orderNo)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while getting order",
		})
		return
	}
	emptyOrder := form.Order{
		Header:  form.OrderHeader{},
		Details: []form.OrderDetail{},
	}
	if reflect.DeepEqual(order, emptyOrder) {
		log.Println("No order belong to this orderNo.")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No order belong to this orderNo.",
		})
		return

	}

	// change order_date to local time
	location, err := time.LoadLocation("Asia/Bangkok")
	order.Header.OrderDate = (order.Header.OrderDate).In(location)

	log.Println("Get order successfully")
	c.JSON(http.StatusOK, order)
}
