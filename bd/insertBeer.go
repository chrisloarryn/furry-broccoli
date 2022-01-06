package bd

import (
	"context"
	"errors"
	"time"

	"github.com/ccontreras/furry-broccoli/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertTweet function inserts a tweet register into the database
func InsertBeer(t models.SaveBeer) (*models.Beer, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var insertedElement models.Beer

	db := MongoCN.Database("furry-broccoli")
	col := db.Collection("beers")


	// TODO: Check if the beer already exists
	beer, _ := FindABeer(t.BeerID);
		
	if beer.BeerID == t.BeerID {
		return nil, false, errors.New("can not be save because beer already exists")
	}	
	


	register := bson.M{
		"beerId": t.BeerID,
		"name": t.Name,
		"brewery": t.Brewery,
		"country": t.Country,
		"price": t.Price,
		"currency": t.Currency,
		"createdAt": time.Now(),
		"updatedAt": time.Now(),
	}

	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return nil, false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	condition := bson.M{
		"_id": objID,
	}
	err = col.FindOne(ctx, condition).Decode(&insertedElement)
	if err != nil {
		return nil, false, err
	}

	return &insertedElement, true, nil
}
