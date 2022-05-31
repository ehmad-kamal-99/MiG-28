//go:generate mockgen -destination=./../mocks/beer_storage.go -package=mocks . BeerStorage
package core

import (
	domain "backend-svc-template"
)

// BeerStorage defines the interface required for handling beer storage related operations.
type BeerStorage interface {
	Add(beer *domain.Beer) (string, error)
	Get(id string) (*domain.Beer, error)
	Edit(beer *domain.Beer) error
	List() ([]*domain.Beer, error)
	Delete(id string) error
}

type beer struct {
	bd BeerStorage
}

func NewBeer(bd BeerStorage) *beer {
	return &beer{bd: bd}
}

func (b *beer) Add(beer *domain.Beer) (*domain.Beer, error) {
	id, err := b.bd.Add(beer)
	if err != nil {
		return nil, err
	}

	beer.ID = id

	return beer, nil
}

func (b *beer) Get(id string) (*domain.Beer, error) {
	return b.bd.Get(id)
}

func (b *beer) Edit(beer *domain.Beer) (*domain.Beer, error) {
	if err := b.bd.Edit(beer); err != nil {
		return nil, err
	}

	return beer, nil
}

func (b *beer) List(filter, sort []string, limit, offset int) ([]*domain.Beer, error) {
	return b.bd.List()
}

func (b *beer) Delete(id string) error {
	return b.bd.Delete(id)
}
