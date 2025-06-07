package domain

// Tablet represents the data structure for a tablet item
type Tablet struct {
	Title       string  `json:"title" pg:"title"`
	Price       float64 `json:"price" pg:"price"`
	Description string  `json:"description" pg:"description"`
	Rating      int     `json:"rating" pg:"rating"`
	ImageURL    string  `json:"image_url" pg:"image_url"`
}
