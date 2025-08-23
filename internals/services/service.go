package services

import "go.mongodb.org/mongo-driver/mongo"

type Storage struct {
	Todo interface {
		Insert(Todo) error
	}
}


func returnCollectionPointer(collection string, client *mongo.Client) *mongo.Collection {
	return client.Database("todos-db").Collection(collection)
}

func NewStorage(mongo *mongo.Client) Storage {
	return Storage{
		Todo: &TodoStore{mongoclient: mongo},
	}
}
