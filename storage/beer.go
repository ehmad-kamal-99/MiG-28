package storage

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type beer struct {
	client *mongo.Client
}

func NewBeer(uri string) *beer {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("failed to connect to mongo cluster, uri: %s, err: %+v", uri, err)
	}

	return &beer{client: client}
}

func (b *beer) Add(beer *backend_svc_template.Beer) (string, error) {
	// TODO: handle logic to add beer
	return "", nil
}

func (b *beer) Get(id string) (*backend_svc_template.Beer, error) {
	// TODO: handle logic to get beer
	return nil, nil
}

func (b *beer) Edit(beer *backend_svc_template.Beer) error {
	// TODO: handle logic to edit beer
	return nil
}

func (b *beer) List() ([]*backend_svc_template.Beer, error) {
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
