package bd

import (
	"context"
	"time"

	"github.com/JuanMira/tweetgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadFollowersTweets(id string, page int) ([]models.RetrieveFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tweetGo")
	col := db.Collection("relations")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"UserID": id}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "UserRelationId",
			"foreignField": "UserId",
			"as":           "tweet",
		},
	})
	// los resultados vienen de forma para poder procesarlos
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.Date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip}) //first skip
	conditions = append(conditions, bson.M{"$limit": 20})  //second limit

	cursor, err := col.Aggregate(ctx, conditions)

	var result []models.RetrieveFollowers

	err = cursor.All(ctx, &result)

	if err != nil {
		return result, false
	}

	return result, true
}
