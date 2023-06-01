package entities

type State struct {
	ID      string `json:"id" db:"id" bson:"_id"`
	Name    string `json:"name" db:"name" bson:"name"`
	Country string `json:"country" db:"country" bson:"country"`
}
