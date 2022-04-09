package bd

import (
	"context"
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

	// mongodb options
	//findOptions := options.FindOne()
	//findOptions.SetProjection(bson.M{
	//	"_id":      1,
	//	"name":     1,
	//	"location": 1,
	//	"price":    1,
	//})
	//"description": 1,
	//"available": 1, "bedrooms": 1, "bathrooms": 1, "maxguests": 1, "owner": 1,
	//"reviews": 1,

	err := col.FindOne(ctx, condition).Decode(&searches)
	//fmt.Printf("Entry %v", condition)

	if err != nil {
		//fmt.Println("Entry not found " + err.Error())
		return searches, err
	}

	return searches, nil

}
