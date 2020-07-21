package entities

// Product represents a single product
type Product struct {
	ID          string  `json:"id" bson:"id"`
	Name        string  `json:"name" bson:"name"`
	Price       float64 `json:"price" bson:"price"`
	Description string  `json:"description" bson:"description"`
	Image       string  `json:"image" bson:"image"`
}
