package endpoints

import (
	"agarwalconsulting.io/rvstore/orders/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

type OrderEndpoints struct {
	Index  endpoint.Endpoint
	Show   endpoint.Endpoint
	Create endpoint.Endpoint
}

// MakeEndpoints initializes the OrderEndpoints for OrderService
func MakeEndpoints(s service.OrdersService) OrderEndpoints {
	return OrderEndpoints{
		Index:  makeIndexEndpoint(s),
		Show:   makeShowEndpoint(s),
		Create: makeCreateEndpoint(s),
	}
}
