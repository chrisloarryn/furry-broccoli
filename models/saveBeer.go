package models

// SaveBeer is the format or structure that our tweet is gonna be
type SaveBeer struct {
	BeerID		 int             		`bson:"beerId" json:"beerId,omitempty"`
	Name       string             `bson:"name" json:"name,omitempty"`
	Brewery    string             `bson:"brewery" json:"brewery,omitempty"`
	Country    string             `bson:"country" json:"country"`
	Price      float64            `bson:"price" json:"price,omitempty"`
	Currency   string             `bson:"currency" json:"currency,omitempty"`
}
