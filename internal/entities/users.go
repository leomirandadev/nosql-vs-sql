package entities

type User struct {
	ID      int    `json:"id" db:"id" bson:"id"`
	Name    string `json:"name" db:"name" bson:"name"`
	City    string `json:"city" db:"city" bson:"city"`
	State   string `json:"state" db:"state" bson:"state"`
	Country string `json:"country" db:"country" bson:"country"`
}
