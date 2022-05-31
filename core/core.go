//go:generate mockgen -destination=./../mocks/storage.go -package=mocks . BeerStorage,ReviewStorage
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
	bs BeerStorage
}

func NewBeer(bs BeerStorage) *beer {
	return &beer{bs: bs}
}

func (b *beer) Add(beer *domain.Beer) (*domain.Beer, error) {
	id, err := b.bs.Add(beer)
	if err != nil {
		return nil, err
	}

	beer.ID = id

	return beer, nil
}

func (b *beer) Get(id string) (*domain.Beer, error) {
	return b.bs.Get(id)
}

func (b *beer) Edit(beer *domain.Beer) (*domain.Beer, error) {
	if err := b.bs.Edit(beer); err != nil {
		return nil, err
	}

	return beer, nil
}

func (b *beer) List(filter, sort []string, limit, offset int) ([]*domain.Beer, error) {
	return b.bs.List()
}

func (b *beer) Delete(id string) error {
	return b.bs.Delete(id)
}

// ReviewStorage defines the interface required for handling review storage related operations.
type ReviewStorage interface {
	Add(review *domain.Review) (string, error)
	Delete(id string) error
	List(beerID string) ([]*domain.Review, error)
}

type review struct {
	rd ReviewStorage
}

func NewReview(rd ReviewStorage) *review {
	return &review{rd: rd}
}

func (rs *review) Add(review *domain.Review) (*domain.Review, error) {
	id, err := rs.rd.Add(review)
	if err != nil {
		return nil, err
	}

	review.ID = id

	return review, nil
}

func (rs *review) Delete(id string) error {
	return rs.rd.Delete(id)
}

func (rs *review) List(beerID string) ([]*domain.Review, error) {
	return rs.rd.List(beerID)
}
