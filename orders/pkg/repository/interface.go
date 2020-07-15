package repository

import "agarwalconsulting.io/rvstore/entities"

type OrdersRepository interface {
	FetchAll() ([]entities.Order, error)
	FindBy(ID string) (entities.Order, error)
	Save(*entities.Order) error
}
