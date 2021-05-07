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
	ImageURL    string  `json:"image_url"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	UnitPrice   float64 `json:"unit_price"`
	PublishYear uint64  `json:"publish_year"`
	Publisher   string  `json:"publisher"`
	Edition     uint64  `json:"edition"`
	Category    string  `json:"category"`
}

type BookRequest struct {
	ISBN        string  `json:"isbn"`
	Name        string  `json:"name"`
	ImageURL    string  `json:"image_url"`
	Author      uint64  `json:"author"`
	Description string  `json:"description"`
	UnitPrice   float64 `json:"unit_price"`
	PublishYear uint64  `json:"publish_year"`
	Publisher   uint64  `json:"publisher"`
	Edition     uint64  `json:"edition"`
	Category    uint64  `json:"category"`
}
