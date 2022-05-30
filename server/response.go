package server

import (
	domain "backend-svc-template"
)

type getBeerRes struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

type editBeerRes struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

type listBeerRes struct {
	Beers []domain.Beer `json:"beers"`
}

type listReviewRes struct {
	Reviews []domain.Review `json:"reviews"`
}
