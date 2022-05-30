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
	List() ([]*domain.Beer, error)
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

}

func (b *beer) edit(ctx *gin.Context) {

}

func (b *beer) list(ctx *gin.Context) {

}

func (b *beer) delete(ctx *gin.Context) {

}
