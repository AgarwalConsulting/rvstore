package endpoints

import (
	"context"

	"agarwalconsulting.io/rvstore/entities"
	"agarwalconsulting.io/rvstore/orders/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

type CreateRequest struct {
	Order entities.Order
}

type CreateResponse struct {
	Order entities.Order
}

func makeCreateEndpoint(s service.OrdersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		order, err := s.Create(ctx, req.Order)
		return CreateResponse{Order: order}, err
	}
}
