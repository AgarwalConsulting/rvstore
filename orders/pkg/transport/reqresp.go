package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"agarwalconsulting.io/rvstore/orders/pkg/endpoints"
	"github.com/gorilla/mux"
)

func decodeIndexRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return endpoints.IndexRequest{}, nil
}

func encodeIndexResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(endpoints.IndexResponse)
	return json.NewEncoder(w).Encode(res.Orders)
}

func decodeShowRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var showReq endpoints.ShowRequest

	vars := mux.Vars(req)
	showReq.ID = vars["id"]

	return showReq, nil
}

func encodeShowResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(endpoints.ShowResponse)
	return json.NewEncoder(w).Encode(res.Order)
}

func decodeCreateRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var createReq endpoints.CreateRequest

	json.NewDecoder(req.Body).Decode(&createReq.Order)

	return createReq, nil
}

func encodeCreateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(endpoints.CreateResponse)
	return json.NewEncoder(w).Encode(res.Order)
}
