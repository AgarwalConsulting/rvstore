package entities

import (
	"time"
)

// Order represents a single order with multiple products
type Order struct {
	ID           string     `json:"id" bson:"id"`
	CreatedAt    *time.Time `json:"orderDate" bson:"orderDate"`
	CustomerName string     `json:"customerName" bson:"customerName"`
	SubTotal     float64    `json:"subTotal" bson:"subTotal"`
	Tax          float64    `json:"tax" bson:"tax"`
	Total        float64    `json:"total" bson:"total"`
	Items        []Product  `json:"items" bson:"items"`
}
