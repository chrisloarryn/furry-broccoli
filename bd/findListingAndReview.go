package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/ccontreras/furry-broccoli/models"
	"go.mongodb.org/mongo-driver/bson"
)

// FindListingAndReviewId is a function to find a beer in the database
func FindListingAndReviewId(searchID string) (models.ListingAndReview, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("sample_airbnb")
	col := db.Collection("listingsAndReviews")

	var searches models.ListingAndReview
	// objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": searchID,
	}
	// "_id": objID,

	err := col.FindOne(ctx, condition).Decode(&searches)
	fmt.Printf("Entry not found %v", condition)

	if err != nil {
		fmt.Println("Entry not found " + err.Error())
		return searches, err
	}

	return searches, nil

}
