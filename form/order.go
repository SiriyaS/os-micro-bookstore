package form

import "time"

type OrderReq struct {
	Header  OrderHeaderReq `json:"header"`
	Details []OrderDetail  `json:"details"`
}

type Order struct {
	Header  OrderHeader   `json:"header"`
	Details []OrderDetail `json:"details"`
}

type OrderHeaderReq struct {
	OrderNo    string    `json:"order_no"`
	User       uint64    `json:"user"`
	OrderDate  time.Time `json:"order_date"`
	GrandTotal float64   `json:"grand_total"`
	Address    string    `json:"address"`
}

type OrderHeader struct {
	OrderNo    string    `json:"order_no"`
	Userame    string    `json:"username"`
	Address    string    `json:"address"`
	OrderDate  time.Time `json:"order_date"`
	GrandTotal float64   `json:"grand_total"`
}

type OrderDetail struct {
	OrderNo   string  `json:"order_no"`
	OrderSeq  uint64  `json:"order_seq"`
	BookISBN  string  `json:"book_isbn"`
	Quantity  uint64  `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	Total     float64 `json:"total"`
}
