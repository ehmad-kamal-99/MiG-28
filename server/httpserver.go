package server

import (
	"net/http"

	"backend-svc-template/middleware"
)

type server struct {
	beerService   BeerService
	reviewService ReviewService

	firebaseAuth *middleware.FirebaseAuth
	mux          *http.ServeMux
}

func New(bs BeerService, rs ReviewService, auth *middleware.FirebaseAuth) *server {
	s := &server{
		beerService:   bs,
		reviewService: rs,
		firebaseAuth:  auth,
	}

	bh := beer{bs: bs}
	rh := review{rs: rs}

	//bg := engine.Group("beer/v1")
	//// bg.Use(auth.Check())
	//{
	//	bg.POST("/beer", bh.add)
	//	bg.GET("/beer", bh.get)
	//	bg.GET("/beers", bh.list)
	//	bg.PUT("/beer", bh.edit)
	//	bg.DELETE("/beer", bh.delete)
	//}
	//
	//rg := engine.Group("review/v1")
	//// rg.Use(auth.Check())
	//{
	//	rg.POST("/review", rh.add)
	//	rg.GET("/review", rh.list)
	//	rg.DELETE("/review", rh.delete)
	//}

	mux := http.NewServeMux()
	mux.HandleFunc("/beer", bh.add)
	mux.HandleFunc("/beer", bh.add)
	mux.HandleFunc("/beers", bh.add)
	mux.HandleFunc("/beer", bh.add)

	return s
}

func (s *server) Mux() *http.ServeMux {
	return s.mux
}
