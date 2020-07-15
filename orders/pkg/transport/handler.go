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

	r.Methods("GET").Path("/orders").Handler(httptransport.NewServer(
		endpoints.Index,
		decodeIndexRequest,
		encodeIndexResponse,
	))

	r.Methods("GET").Path("/orders/{id}").Handler(httptransport.NewServer(
		endpoints.Show,
		decodeShowRequest,
		encodeShowResponse,
	))

	r.Methods("POST").Path("/orders").Handler(httptransport.NewServer(
		endpoints.Create,
		decodeCreateRequest,
		encodeCreateResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
