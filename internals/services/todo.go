package services

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type TodoStore struct {
	mongoclient *mongo.Client
}

type Todo struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string    `json:"task" bson:"task"`
	Completed bool      `json:"completed" bson:"completed"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

func (store *TodoStore) Insert(todo Todo) error {
	collection := returnCollectionPointer("todos", store.mongoclient)
	_, err := collection.InsertOne(context.TODO(), Todo{
		Task:      todo.Task,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	)

	return err
}

func (store *TodoStore) Update(Todo) error {
	return nil
}

func (store *TodoStore) GetByID(id string) (*Todo, error) {
	return nil, nil
}

func (store *TodoStore) GetAllForUser(userID string) ([]*Todo, error) {
	return nil, nil
}

func (store *TodoStore) Delete(id string) error {
	return nil
}
