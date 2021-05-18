package bd

import (
	"context"
	"time"

	"github.com/JuanMira/tweetgo/models"
)

func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 14*time.Second)
	defer cancel()

	db := MongoCN.Database("tweetGo")
	col := db.Collection("relations")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}

	return true, nil
}
