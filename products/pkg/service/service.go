package service

import (
	"context"

	"agarwalconsulting.io/rvstore/entities"
	"agarwalconsulting.io/rvstore/products/pkg/repository"
)

// ProductsService describes the service.
type ProductsService interface {
	Index(ctx context.Context) (products []entities.Product, err error)
	IndexSecure(ctx context.Context) (products []entities.Product, err error)
}

type basicProductsService struct {
	repo repository.ProductsRepository
}

func (b *basicProductsService) Index(ctx context.Context) (products []entities.Product, err error) {
	return b.repo.FetchAll()
}

// NewBasicProductsService returns a naive, stateless implementation of ProductsService.
func NewBasicProductsService(repo repository.ProductsRepository) ProductsService {
	return &basicProductsService{repo}
}

// New returns a ProductsService with all of the expected middleware wired in.
func New(repo repository.ProductsRepository, middleware []Middleware) ProductsService {
	var svc ProductsService = NewBasicProductsService(repo)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicProductsService) IndexSecure(ctx context.Context) (products []entities.Product, err error) {
	return b.repo.FetchAll()
}
