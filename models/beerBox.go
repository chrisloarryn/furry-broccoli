package models

// Beer data structure for save the beer information
type BeerBox struct {
	TotalPrice       float64             `bson:"totalPrice" json:"totalPrice,omitempty"`
}
