package bd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chrisloarryn/furry-broccoli/models"
	"go.mongodb.org/mongo-driver/bson"
)

// SearchBeers function for reading tweets
func SearchBeers() ([]*models.Beer, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	fmt.Println("Searching beers")

	db := MongoCN.Database("furry-broccoli")
	col := db.Collection("beers")

	var results []*models.Beer

	findCondition := bson.M{}

	cursor, err := col.Find(ctx, findCondition)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var register models.Beer
		err := cursor.Decode(&register)
		if err != nil {
			return results, false
		}
		results = append(results, &register)
	}

	return results, true
}
