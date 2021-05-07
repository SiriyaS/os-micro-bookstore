package form

type Product struct {
	ProductID       uint64  `json:"product_id"`
	ProductName     string  `json:"product_name"`
	ProductQuantity uint64  `json:"product_quantity"`
	ProductPrice    float64 `json:"product_price"`
}

type Book struct {
	ISBN        string  `json:"isbn"`
	Name        string  `json:"name"`
	Author      string  `json:"author"`
	UnitPrice   float64 `json:"unit_price"`
	PublishYear uint64  `json:"publish_year"`
	Publisher   string  `json:"publisher"`
	Edition     uint64  `json:"edition"`
	Category    string  `json:"category"`
	Image_URL   string  `json:"image_url"`
}

type BookRequest struct {
	ISBN        string  `json:"isbn"`
	Name        string  `json:"name"`
	Author      uint64  `json:"author"`
	UnitPrice   float64 `json:"unit_price"`
	PublishYear uint64  `json:"publish_year"`
	Publisher   uint64  `json:"publisher"`
	Edition     uint64  `json:"edition"`
	Category    uint64  `json:"category"`
	Image_URL   string  `json:"image_url"`
}
