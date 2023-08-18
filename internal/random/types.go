package random

type Item struct {
		ID          int                `json:"id"`
		Title       string             `json:"title"`
		Price       float32            `json:"price"`
		Category    string             `json:"category"`
		Description string             `json:"description"`
		Image       string             `json:"image"`
		Rating      map[string]float64 `json:"rating"`
	}