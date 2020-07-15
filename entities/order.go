package entities

import (
	"time"
)

// Order represents a single order with multiple products
type Order struct {
	ID           string     `json:"id"`
	CreatedAt    *time.Time `json:"orderDate"`
	CustomerName string     `json:"customerName"`
	SubTotal     float64    `json:"subTotal"`
	Tax          float64    `json:"tax"`
	Total        float64    `json:"total"`
	Items        []Product  `json:"items"`
}
