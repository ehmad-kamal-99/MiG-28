package server

import (
	"net/http"

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
	req := new(newReviewReq)

	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)

		return
	}

	review, err := r.rs.Add(&domain.Review{
		BeerID:  req.BeerID,
		Comment: req.Comment,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)

		return
	}

	ctx.JSON(http.StatusCreated, review)
}

func (r *review) list(ctx *gin.Context) {
	req := new(listReviewReq)

	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)

		return
	}

	reviews, err := r.rs.List(req.BeerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)

		return
	}

	ctx.JSON(http.StatusOK, reviews)
}

func (r *review) delete(ctx *gin.Context) {
	req := new(deleteReviewReq)

	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)

		return
	}

	if err := r.rs.Delete(req.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)

		return
	}

	ctx.JSON(http.StatusOK, nil)
}
