package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/ccontreras/furry-broccoli/models"
	"go.mongodb.org/mongo-driver/bson"
)

// FindABeer is a function to find a beer in the database
func FindABeer(beerID int) (models.Beer, error) {
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

	if err != nil {
		fmt.Println("Entry not found " + err.Error())
		return beer, err
	}

	return beer, nil

}
