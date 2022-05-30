package server

type newBeerReq struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

type getBeerReq struct {
	ID string `form:"id"`
}

type editBeerReq struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

type listBeerReq struct {
	Filter []string `form:"filter"`
	Sort   []string `form:"sort"`
	Page   int      `form:"page"`
}

type deleteBeerReq struct {
	ID string `form:"id"`
}

type newReviewReq struct {
	BeerID  string `json:"beer_id"`
	Comment string `json:"comment"`
}

type listReviewReq struct {
	BeerID string `form:"beer_id"`
}

type deleteReviewReq struct {
	ID string `form:"id"`
}
