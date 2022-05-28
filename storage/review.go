package storage

import (
	"backend-svc-template"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type review struct {
	client *mongo.Client
}

func NewReview(uri string) *review {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("failed to connect to mongo cluster, uri: %s, err: %+v", uri, err)
	}

	return &review{client: client}
}

func (r *review) Add(review *domain.Review) (string, error) {
	// TODO: handle logic to add review
	return "", nil
}

func (r *review) Delete(id string) error {
	// TODO: handle logic to delete review
	return nil
}

func (r *review) List(beerID string) ([]*domain.Review, error) {
	// TODO: handle logic to list reviews
	return nil, nil
}
