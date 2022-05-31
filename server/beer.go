package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend-svc-template"
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

func (b *beer) add(ctx *gin.Context) {
	req := new(newBeerReq)

	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)

		return
	}

	beer, err := b.bs.Add(&domain.Beer{
		Name:  req.Name,
		Brand: req.Brand,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)

		return
	}

	ctx.JSON(http.StatusCreated, beer)
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
