package main

import (
	"context"
	"net/http"
	"os"

	"agarwalconsulting.io/rvstore/orders/pkg/endpoints"
	"agarwalconsulting.io/rvstore/orders/pkg/repository"
	"agarwalconsulting.io/rvstore/orders/pkg/service"
	"agarwalconsulting.io/rvstore/orders/pkg/transport"
)

var (
	mongoURL string
	httpAddr string
)

func init() {
	var ok bool
	mongoURL, ok = os.LookupEnv("MONGO_DB_URL")
	if !ok {
		mongoURL = "mongodb://root:example@localhost/admin"
	}

	httpAddr, ok = os.LookupEnv("ORDER_SERVICE_ADDR")
	if !ok {
		httpAddr = ":8080"
	}
}

func main() {
	ctx, _ := context.WithCancel(context.Background())

	orderRepo := repository.New(mongoURL)
	orderSvc := service.New(orderRepo)
	orderEndpoints := endpoints.MakeEndpoints(orderSvc)

	r := transport.NewHTTPServer(ctx, orderEndpoints)

	http.ListenAndServe(httpAddr, r)
}
