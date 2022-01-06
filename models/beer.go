package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Beer data structure for save the beer information
type Beer struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	BeerID		  int             		`bson:"beerId" json:"beerId,omitempty"`
	Name        string              `bson:"name" json:"name,omitempty"`
	Brewery     string              `bson:"brewery" json:"brewery,omitempty"`
	Country     string              `bson:"country" json:"country"`
	Price       float64             `bson:"price" json:"price,omitempty"`
	Currency    string              `bson:"currency" json:"currency,omitempty"`
	CreatedAt   time.Time 			  	`bson:"createdAt,omitempty" json:"createdAt,omitempty"` 
	UpdatedAt   time.Time 				  `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
