package entities

type UserDetails struct {
	ID      string `json:"id" db:"id" bson:"_id"`
	Name    string `json:"name" db:"name" bson:"name"`
	City    string `json:"city" db:"city" bson:"city"`
	State   string `json:"state" db:"state" bson:"state"`
	Country string `json:"country" db:"country" bson:"country"`
}
