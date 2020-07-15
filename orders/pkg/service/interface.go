package service

import (
	"context"

	"agarwalconsulting.io/rvstore/entities"
)

// OrdersService describes the service.
type OrdersService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Index(ctx context.Context) ([]entities.Order, error)
	Show(ctx context.Context, id string) (entities.Order, error)
	Create(ctx context.Context, order entities.Order) (entities.Order, error)
}
