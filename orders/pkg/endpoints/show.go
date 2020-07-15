package endpoints

import (
	"context"

	"agarwalconsulting.io/rvstore/entities"
	"agarwalconsulting.io/rvstore/orders/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

type ShowRequest struct {
	ID string
}

type ShowResponse struct {
	Order entities.Order
}

func makeShowEndpoint(s service.OrdersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ShowRequest)
		order, err := s.Show(ctx, req.ID)
		return ShowResponse{Order: order}, err
	}
}
