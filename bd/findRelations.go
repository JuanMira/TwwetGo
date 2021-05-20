package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/JuanMira/tweetgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func FindRelations(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tweetGo")
	col := db.Collection("relations")

	condition := bson.M{
		"UserID":         t.UserId,
		"UserRelationId": t.UserRelationId,
	}

	var result models.Relation

	fmt.Println(result)

	err := col.FindOne(ctx, condition).Decode(&result)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil

}
