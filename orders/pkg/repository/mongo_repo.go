package repository

import (
	"context"
	"errors"
	"time"

	"github.com/gofrs/uuid"

	"agarwalconsulting.io/rvstore/entities"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// OrdersCollection is the name of the collection in mongoDB
const OrdersCollection = "orders"

const databaseName = "rvstore"

var u1 = uuid.Must(uuid.NewV4())

type mongoRepository struct {
	*mongo.Collection
}

func (mr *mongoRepository) FetchAll() ([]entities.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := mr.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var orders []entities.Order

	for cur.Next(ctx) {
		var nextOrder entities.Order
		err := cur.Decode(&nextOrder)
		if err != nil {
			return nil, err
		}
		orders = append(orders, nextOrder)
	}

	return orders, nil
}

func (mr *mongoRepository) FindBy(ID string) (entities.Order, error) {
	filter := bson.M{"id": ID}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var order entities.Order

	err := mr.FindOne(ctx, filter).Decode(&order)

	if err != nil {
		return order, err
	}

	return order, nil
}

func (mr *mongoRepository) Save(order *entities.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	u1, err := uuid.NewV4()

	if err != nil {
		return err
	}

	if order.ID == "" {
		order.ID = u1.String()
	}

	res, err := mr.InsertOne(ctx, order)

	if err != nil {
		return err
	}

	_, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return errors.New("not a primitive.ObjectID type")
	}

	// order.ID = id.Hex()
	return nil
}

// New returns an implementation of OrdersRepository which interacts to mongoDB
func New(ctx context.Context, mongoURL string) (OrdersRepository, error) {
	log.Infof("Connecting to mongo on: %s...\n", mongoURL)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))

	if err != nil {
		return nil, err
	}

	log.Info("Pinging...")
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return nil, err
	}

	log.Infof("Getting the collection %s...\n", OrdersCollection)
	collection := client.Database(databaseName).Collection(OrdersCollection)

	return &mongoRepository{collection}, nil
}
