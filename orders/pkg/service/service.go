package service

import (
	"context"

	"agarwalconsulting.io/rvstore/entities"
	"agarwalconsulting.io/rvstore/orders/pkg/repository"
)

type ordersService struct {
	repo repository.OrdersRepository
}

func (ordSvc *ordersService) Index(ctx context.Context) ([]entities.Order, error) {
	return ordSvc.repo.FetchAll()
}

func (ordSvc *ordersService) Show(ctx context.Context, id string) (entities.Order, error) {
	return ordSvc.repo.FindBy(id)
}

func (ordSvc *ordersService) Create(ctx context.Context, newOrder entities.Order) (entities.Order, error) {
	err := ordSvc.repo.Save(&newOrder)

	return newOrder, err
}

// New creates an instance of OrdersService
func New(repo repository.OrdersRepository) OrdersService {
	return &ordersService{repo}
}
