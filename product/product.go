package product

type Product struct {
	Site     string  `json:"site"`
	NumPrice float64 `json:"float_price"`
	Price    string  `json:"price"`
	Title    string  `json:"title"`
	URL      string  `json:"url"`
	Image    string  `json:"image"`
}
