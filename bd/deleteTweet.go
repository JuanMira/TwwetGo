package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweet(id string, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tweetGo")
	col := db.Collection("tweet")

	objId, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{
		"_id":    objId,
		"UserId": userId,
	}

	_, err := col.DeleteOne(ctx, condition)

	return err
}
