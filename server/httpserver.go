package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend-svc-template/middleware"
)

type server struct {
	beerService   BeerService
	reviewService ReviewService

	firebaseAuth *middleware.FirebaseAuth

	engine *gin.Engine
}

func New(bs BeerService, rs ReviewService, auth *middleware.FirebaseAuth) *server {
	s := &server{
		beerService:   bs,
		reviewService: rs,
		firebaseAuth:  auth,
	}

	engine := gin.New()

	engine.Use(middleware.Recovery())

	// service health point
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})

	bh := beer{bs: bs}
	rh := review{rs: rs}

	bg := engine.Group("beer/v1")
	bg.Use(auth.Check())
	{
		bg.POST("/beer", bh.add)
		bg.GET("/beer", bh.get)
		bg.GET("/beers", bh.list)
		bg.PUT("/beer", bh.edit)
		bg.DELETE("/beer", bh.delete)
	}

	rg := engine.Group("review/v1")
	rg.Use(auth.Check())
	{
		rg.POST("/review", rh.add)
		rg.POST("/review", rh.list)
		rg.POST("/review", rh.delete)
	}

	s.engine = engine

	return s
}

func (s *server) Engine() *gin.Engine {
	return s.engine
}
