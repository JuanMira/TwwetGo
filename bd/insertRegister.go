package bd

import (
	"context"
	"time"

	"github.com/JuanMira/tweetgo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert register in db
func InsertData(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// se ejecuta como ultima instruccion

	// da de baja el context
	defer cancel()

	db := MongoCN.Database("tweetGo")
	col := db.Collection("Users")

	u.Password, _ = EncryptPassword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
