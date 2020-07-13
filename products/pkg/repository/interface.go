package repository

import "agarwalconsulting.io/rvstore/entities"

// ProductsRepository describes a repository for Product
type ProductsRepository interface {
	FetchAll() ([]entities.Product, error)
}
