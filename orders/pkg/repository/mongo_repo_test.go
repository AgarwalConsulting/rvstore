package repository_test

import (
	"context"
	"log"
	"os"
	"testing"

	"agarwalconsulting.io/rvstore/entities"
	"agarwalconsulting.io/rvstore/orders/pkg/repository"
	"github.com/stretchr/testify/assert"
)

var orderRepo repository.OrdersRepository
var mongoURL string

func TestMain(m *testing.M) {
	ctx, cancelFn := context.WithCancel(context.Background())

	url, ok := os.LookupEnv("MONGO_DB_URL")

	if !ok {
		url = "mongodb://root:example@localhost/admin"
	}

	mongoURL = url

	var err error
	orderRepo, err = repository.New(ctx, mongoURL)

	if err != nil {
		log.Fatal(err)
	}

	code := m.Run()
	cancelFn()

	os.Exit(code)
}

func TestSaveAndFindBy(t *testing.T) {
	orderID := "3"
	order := entities.Order{
		ID:           orderID,
		CustomerName: "John Doe",
		SubTotal:     41,
		Tax:          1.12,
		Total:        42.12,
		Items: []entities.Product{
			{ID: "1", Name: "Screw Driver 3000"},
			{ID: "2", Name: "cat5 Cables"},
		},
	}

	err := orderRepo.Save(&order)

	assert.Nil(t, err)

	retrievedOrder, err := orderRepo.FindBy(order.ID)

	assert.Nil(t, err)
	assert.Equal(t, orderID, retrievedOrder.ID)
	assert.Equal(t, order.CustomerName, retrievedOrder.CustomerName)
	assert.NotNil(t, retrievedOrder.CreatedAt)

	allOrders, err := orderRepo.FetchAll()

	t.Logf("found %d orders\n", len(allOrders))

	assert.Nil(t, err)
	if len(allOrders) < 1 {
		t.Fatal("unable to find any orders")
	}
}
