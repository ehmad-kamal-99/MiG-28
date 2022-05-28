package domain

type Beer struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Brand string `json:"brand" bson:"brand"`
}
