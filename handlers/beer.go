package handlers

import (
	"github.com/gin-gonic/gin"
)

type BeerService interface {
	Add(beer *backend_svc_template.Beer) (*backend_svc_template.Beer, error)
	Get(id string) (*backend_svc_template.Beer, error)
	Edit(beer *backend_svc_template.Beer) (*backend_svc_template.Beer, error)
	List() ([]*backend_svc_template.Beer, error)
	Delete(id string) error
}

type beer struct {
	bs BeerService
}

func NewBeer(bs BeerService) *beer {
	return &beer{bs: bs}
}

func (b *beer) Add(ctx *gin.Context) {

}

func (b *beer) Get(ctx *gin.Context) {

}

func (b *beer) Edit(ctx *gin.Context) {

}

func (b *beer) List(ctx *gin.Context) {

}

func (b *beer) Delete(ctx *gin.Context) {

}
