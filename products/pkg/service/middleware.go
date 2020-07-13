package service

import (
	"context"

	"agarwalconsulting.io/rvstore/entities"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(ProductsService) ProductsService

type loggingMiddleware struct {
	logger log.Logger
	next   ProductsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a ProductsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next ProductsService) ProductsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Index(ctx context.Context) (products []entities.Product, err error) {
	defer func() {
		l.logger.Log("method", "Index", "products", products, "err", err)
	}()
	return l.next.Index(ctx)
}

func (l loggingMiddleware) IndexSecure(ctx context.Context) (products []entities.Product, err error) {
	defer func() {
		l.logger.Log("method", "IndexSecure", "products", products, "err", err)
	}()
	return l.next.IndexSecure(ctx)
}
