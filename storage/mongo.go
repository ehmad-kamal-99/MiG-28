package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	domain "backend-svc-template"
)

type beer struct {
	client *mongo.Client
}

func NewBeer(client *mongo.Client) *beer {
	return &beer{client: client}
}

func (b *beer) Add(beer *domain.Beer) (string, error) {
	// TODO: handle logic to add beer
	return "", nil
}

func (b *beer) Get(id string) (*domain.Beer, error) {
	// TODO: handle logic to get beer
	return nil, nil
}

func (b *beer) Edit(beer *domain.Beer) error {
	// TODO: handle logic to edit beer
	return nil
}

func (b *beer) List() ([]*domain.Beer, error) {
	// TODO: handle logic to list beer
	return nil, nil
}

func (b *beer) Delete(id string) error {
	// TODO: handle logic to delete beer
	return nil
}

func (b *beer) Close(ctx context.Context) error {
	return b.client.Disconnect(ctx)
}

type review struct {
	client *mongo.Client
}

func NewReview(client *mongo.Client) *review {
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
