package types

import (
	"fmt"
	"net/http"
)

type Item struct {
	ID          int                `json:"id"`
	Title       string             `json:"title"`
	Price       float32            `json:"price"`
	Category    string             `json:"category"`
	Description string             `json:"description"`
	Image       string             `json:"image"`
	Rating      map[string]float64 `json:"rating"`
}

type CustomHandler struct {
	Message string
}

func (h *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, h.Message)
}
