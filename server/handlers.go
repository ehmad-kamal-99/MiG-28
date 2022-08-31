//go:generate mockgen -destination=./../mocks/service.go -package=mocks . BeerService,ReviewService
package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	domain "backend-svc-template"
)

type BeerService interface {
	Add(beer *domain.Beer) (*domain.Beer, error)
	Get(id string) (*domain.Beer, error)
	Edit(beer *domain.Beer) (*domain.Beer, error)
	List(filter, sort []string, limit, offset int) ([]*domain.Beer, error)
	Delete(id string) error
}

type beer struct {
	bs BeerService
}

func (b *beer) add(w http.ResponseWriter, r *http.Request) {
	var br newBeerReq

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		if _, err := w.Write([]byte("failed to read body")); err != nil {
			log.Fatal("failed to write response")
		}
	}

	if err := json.Unmarshal(body, &br); err != nil {
		w.WriteHeader(400)
	}

	beer, err := b.bs.Add(&domain.Beer{
		Name:  br.Name,
		Brand: br.Brand,
	})
	if err != nil {
		w.WriteHeader(500)
		if _, err := w.Write([]byte("failed to read body")); err != nil {
			log.Fatal("failed to write response")
		}

		return
	}

	resp, err := json.Marshal(beer)
	if err != nil {
		w.WriteHeader(500)
		if _, err := w.Write([]byte("failed to marshal response")); err != nil {
			log.Fatal("failed to write response")
		}

		return
	}

	if _, err := w.Write(resp); err != nil {
		w.WriteHeader(500)
		if _, err := w.Write([]byte("failed to write response")); err != nil {
			log.Fatal("failed to write response")
		}

		return
	}

	w.WriteHeader(200)
}

func (b *beer) get(ctx *gin.Context) {
	req := new(getBeerReq)

	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)

		return
	}

	beer, err := b.bs.Get(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)

		return
	}

	ctx.JSON(http.StatusOK, beer)

}

func (b *beer) edit(ctx *gin.Context) {
	req := new(editBeerReq)

	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)

		return
	}

	beer, err := b.bs.Edit(&domain.Beer{
		ID:    req.ID,
		Name:  req.Name,
		Brand: req.Brand,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)

		return
	}

	ctx.JSON(http.StatusOK, beer)

}

func (b *beer) list(ctx *gin.Context) {
	req := new(listBeerReq)

	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)

		return
	}

	beers, err := b.bs.List(req.Filter, req.Sort, req.Limit, req.Offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)

		return
	}

	ctx.JSON(http.StatusOK, beers)
}

func (b *beer) delete(ctx *gin.Context) {
	req := new(deleteBeerReq)

	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)

		return
	}

	if err := b.bs.Delete(req.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)

		return
	}

	ctx.JSON(http.StatusOK, nil)
}

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
