package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	endpoint "agarwalconsulting.io/rvstore/products/pkg/endpoint"

	http1 "github.com/go-kit/kit/transport/http"
)

// makeIndexHandler creates the handler logic
func makeIndexHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	productsIndexHandler := http1.NewServer(endpoints.IndexEndpoint, decodeIndexRequest, encodeIndexResponse, options...)
	m.Handle("/products", productsIndexHandler)
	m.Handle("/products/", productsIndexHandler)
}

// decodeIndexRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeIndexRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.IndexRequest{}
	return req, nil
}

// encodeIndexResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeIndexResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res, _ := response.(endpoint.IndexResponse)
	err = json.NewEncoder(w).Encode(res.Products)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

// makeIndexSecureHandler creates the handler logic
func makeIndexSecureHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	indexSecureHandler := http1.NewServer(endpoints.IndexSecureEndpoint, decodeIndexSecureRequest, encodeIndexSecureResponse, options...)

	m.Handle("/products/secure", indexSecureHandler)
	m.Handle("/products/secure/", indexSecureHandler)
}

// decodeIndexSecureRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeIndexSecureRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.IndexSecureRequest{}
	return req, nil
}

// encodeIndexSecureResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeIndexSecureResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res, _ := response.(endpoint.IndexSecureResponse)
	err = json.NewEncoder(w).Encode(res.Products)
	return
}
