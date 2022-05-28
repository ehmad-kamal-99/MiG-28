package core

import (
	"backend-svc-template"
)

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
