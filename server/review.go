package server

import (
	"github.com/gin-gonic/gin"

	"backend-svc-template"
)

type ReviewService interface {
	Add(review *domain.Review) (*domain.Review, error)
	Delete(id string) error
	List(beerID string) ([]*domain.Review, error)
}

type review struct {
	rs ReviewService
}

func (r *review) add(ctx *gin.Context) {

}

func (r *review) delete(ctx *gin.Context) {

}

func (r *review) list(ctx *gin.Context) {

}
