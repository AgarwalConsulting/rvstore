package repository

import "agarwalconsulting.io/rvstore/entities"

// TODO: Implement the interaction with mongoDB
type mongoRepository struct {
}

func (mr *mongoRepository) FetchAll() ([]entities.Order, error) {
	return nil, nil
}

func (mr *mongoRepository) FindBy(ID string) (entities.Order, error) {
	return entities.Order{}, nil
}

func (mr *mongoRepository) Save(*entities.Order) error {
	return nil
}

// New returns an implementation of OrdersRepository which interacts to mongoDB
func New(mongoURL string) OrdersRepository {
	return &mongoRepository{}
}
