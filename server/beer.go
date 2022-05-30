package server

import (
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

}

func (b *beer) get(ctx *gin.Context) {

}

func (b *beer) edit(ctx *gin.Context) {

}

func (b *beer) list(ctx *gin.Context) {

}

func (b *beer) delete(ctx *gin.Context) {

}
