package transport

import (
	"context"
	"net/http"

	"agarwalconsulting.io/rvstore/orders/pkg/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints endpoints.OrderEndpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	ordersIndexHandler := httptransport.NewServer(
		endpoints.Index,
		decodeIndexRequest,
		encodeIndexResponse,
	)
	r.Methods("GET").Path("/orders").Handler(ordersIndexHandler)
	r.Methods("GET").Path("/orders/").Handler(ordersIndexHandler)

	orderShowHandler := httptransport.NewServer(
		endpoints.Show,
		decodeShowRequest,
		encodeShowResponse,
	)
	r.Methods("GET").Path("/orders/{id}").Handler(orderShowHandler)
	r.Methods("GET").Path("/orders/{id}/").Handler(orderShowHandler)

	orderCreateHandler := httptransport.NewServer(
		endpoints.Create,
		decodeCreateRequest,
		encodeCreateResponse,
	)

	r.Methods("POST").Path("/orders").Handler(orderCreateHandler)
	r.Methods("POST").Path("/orders/").Handler(orderCreateHandler)

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
