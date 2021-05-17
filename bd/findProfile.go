package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/JuanMira/tweetgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindProfile(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("tweetGo")
	col := db.Collection("Users")

	var profile models.User
	objId, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{
		"_id": objId,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)

	profile.Password = ""

	if err != nil {
		fmt.Println("Row not found " + err.Error())
		return profile, err
	}

	return profile, nil
}
