package bd

import (
	"context"
	"time"

	"github.com/JuanMira/tweetgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckUserExist(e string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	db := MongoCN.Database("tweetGo")
	col := db.Collection("Users")
	defer cancel()
	condition := bson.M{"email": e}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)

	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
