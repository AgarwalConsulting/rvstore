package endpoints

import (
	"context"

	"agarwalconsulting.io/rvstore/entities"
	"agarwalconsulting.io/rvstore/orders/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

type IndexRequest struct{}

type IndexResponse struct {
	Orders []entities.Order
}

func makeIndexEndpoint(s service.OrdersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		orders, err := s.Index(ctx)
		return IndexResponse{Orders: orders}, err
	}
}
