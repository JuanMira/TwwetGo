package bd

import (
	"context"
	"log"
	"time"

	"github.com/JuanMira/tweetgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RetrieveTweets(Id string, page int64) ([]*models.RetrieveTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tweetGo")
	col := db.Collection("tweet")

	var results []*models.RetrieveTweet

	condition := bson.M{
		"UserId": Id,
	}

	options := options.Find()
	//pagination with mongo
	//give the date desc
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "Date", Value: -1}})
	options.SetSkip((page - 1) * 20)

	//pointer
	cursor, err := col.Find(ctx, condition, options)

	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	//todo create a new context
	for cursor.Next(context.TODO()) {
		var register models.RetrieveTweet
		err := cursor.Decode(&register)

		if err != nil {
			return results, false
		}

		results = append(results, &register)
	}
	return results, true
}
