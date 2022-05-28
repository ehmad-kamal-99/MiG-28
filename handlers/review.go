package handlers

import (
	"backend-svc-template"
	"github.com/gin-gonic/gin"
)

type ReviewService interface {
	Add(review *domain.Review) (*domain.Review, error)
	Delete(id string) error
	List(beerID string) ([]*domain.Review, error)
}

type review struct {
	rs ReviewService
}

func NewReview(rs ReviewService) *review {
	return &review{rs: rs}
}

func (r *review) Add(ctx *gin.Context) {

}

func (r *review) Delete(ctx *gin.Context) {

}

func (r *review) List(ctx *gin.Context) {

}
