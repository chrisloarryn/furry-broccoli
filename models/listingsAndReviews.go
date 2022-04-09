package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ListingAndReview struct {
	_id            string               `json:"_id" bson:"_id"`
	Name           string               `json:"name" bson:"name"`
	Description    string               `json:"description" bson:"description"`
	Bed_type       string               `json:"bed_type" bson:"bed_type"`
	Room_type      string               `json:"room_type" bson:"room_type"`
	Summary        string               `json:"summary" bson:"summary"`
	Price          primitive.Decimal128 `json:"price" bson:"price"`
	Minimum_nights string               `json:"minimum_nights" bson:"minimum_nights"`
	Maximum_nights string               `json:"maximum_nights" bson:"maximum_nights"`
}
