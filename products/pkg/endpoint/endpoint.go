package endpoint

import (
	"context"
	service "agarwalconsulting.io/rvstore/products/pkg/service"

	"agarwalconsulting.io/rvstore/entities"
	endpoint "github.com/go-kit/kit/endpoint"
)

// IndexRequest collects the request parameters for the Index method.
type IndexRequest struct{}

// IndexResponse collects the response parameters for the Index method.
type IndexResponse struct {
	Products []entities.Product `json:"products"`
	Err      error              `json:"err"`
}

// MakeIndexEndpoint returns an endpoint that invokes Index on the service.
func MakeIndexEndpoint(s service.ProductsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := s.Index(ctx)
		return IndexResponse{
			Err:      err,
			Products: products,
		}, nil
	}
}

// Failed implements Failer.
func (r IndexResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Index implements Service. Primarily useful in a client.
func (e Endpoints) Index(ctx context.Context) (products []entities.Product, err error) {
	request := IndexRequest{}
	response, err := e.IndexEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(IndexResponse).Products, response.(IndexResponse).Err
}

// IndexSecureRequest collects the request parameters for the IndexSecure method.
type IndexSecureRequest struct{}

// IndexSecureResponse collects the response parameters for the IndexSecure method.
type IndexSecureResponse struct {
	Products []entities.Product `json:"products"`
	Err      error              `json:"err"`
}

// MakeIndexSecureEndpoint returns an endpoint that invokes IndexSecure on the service.
func MakeIndexSecureEndpoint(s service.ProductsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := s.IndexSecure(ctx)
		return IndexSecureResponse{
			Err:      err,
			Products: products,
		}, nil
	}
}

// Failed implements Failer.
func (r IndexSecureResponse) Failed() error {
	return r.Err
}

// IndexSecure implements Service. Primarily useful in a client.
func (e Endpoints) IndexSecure(ctx context.Context) (products []entities.Product, err error) {
	request := IndexSecureRequest{}
	response, err := e.IndexSecureEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(IndexSecureResponse).Products, response.(IndexSecureResponse).Err
}
