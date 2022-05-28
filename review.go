package domain

type Review struct {
	ID      string `json:"id" bson:"_id"`
	BeerID  string `json:"beer_id" bson:"beer_id"`
	Comment string `json:"comment" bson:"comment"`
	Owner   User   `json:"owner" bson:"owner"`
}
