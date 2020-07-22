package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"agarwalconsulting.io/rvstore/orders/pkg/endpoints"
	"agarwalconsulting.io/rvstore/orders/pkg/repository"
	"agarwalconsulting.io/rvstore/orders/pkg/service"
	"agarwalconsulting.io/rvstore/orders/pkg/transport"
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
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

	mongoURL = strings.Trim(mongoURL, "\n")

	httpAddr, ok = os.LookupEnv("ORDER_SERVICE_ADDR")
	if !ok {
		httpAddr = ":9002"
	}
}

func main() {
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	orderRepo, err := repository.New(ctx, mongoURL)

	if err != nil {
		log.Fatal(err)
	}

	orderSvc := service.New(orderRepo)
	orderEndpoints := endpoints.MakeEndpoints(orderSvc)

	r := transport.NewHTTPServer(ctx, orderEndpoints)

	loggedRouter := handlers.LoggingHandler(os.Stdout, handlers.CORS()(r))

	log.Infof("Starting orders service on: %s...\n", httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, loggedRouter))
}
