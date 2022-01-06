package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/ccontreras/furry-broccoli/models"
	"go.mongodb.org/mongo-driver/bson"
)

// FindABeer is a function to find a beer in the database
func BoxBeerPriceByBeerId(beerID int, currency string, quantity int) (models.BeerBox, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("furry-broccoli")
	col := db.Collection("beers")

	var beer models.Beer
	// objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"beerId": beerID,
	}
	// "_id": objID,

	err := col.FindOne(ctx, condition).Decode(&beer)

	defaultPrice := models.BeerBox{
		TotalPrice: 0,
	}

	if err != nil {
		fmt.Println("Entry not found " + err.Error())
		return defaultPrice, err
	}

	// mixedInt64PriceVar := int64(beer.Price)
	mixedFloat64QuantityVar := float64(quantity)

	defaultPrice.TotalPrice = beer.Price * mixedFloat64QuantityVar
	return defaultPrice, nil

}
