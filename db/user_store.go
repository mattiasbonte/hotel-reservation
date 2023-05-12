package db

import (
	"github.com/fulltimegodev/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	GetUserByID(string) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
	}
}

func (s *MongoUserStore) GetUserByID(id string) (*types.User, error) {
    

}
