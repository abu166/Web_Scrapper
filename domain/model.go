package domain

// Tablet represents the data structure for a tablet item
type Tablet struct {
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Rating      int     `json:"rating"`
	ImageURL    string  `json:"image_url"`
}