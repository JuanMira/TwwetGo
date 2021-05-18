package bd

import (
	"context"
	"time"

	"github.com/JuanMira/tweetgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//function to modify profile
func ModifyProfile(u models.User, id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tweetGo")
	col := db.Collection("Users")

	// crea slices o mapas
	registro := make(map[string]interface{})

	if len(u.Name) > 0 {
		registro["Name"] = u.Name
	}

	if len(u.LastName) > 0 {
		registro["LastName"] = u.LastName
	}

	if len(u.Email) > 0 {
		registro["Email"] = u.Email
	}

	registro["BirthDate"] = u.BirthDate

	if len(u.Avatar) > 0 {
		registro["Avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		registro["Banner"] = u.Banner
	}

	if len(u.WebSite) > 0 {
		registro["WebSite"] = u.WebSite
	}

	if len(u.Location) > 0 {
		registro["Location"] = u.Location
	}

	updtString := bson.M{
		"$set": registro,
	}

	objId, _ := primitive.ObjectIDFromHex(id)
	//filter
	filter := bson.M{"_id": bson.M{"$eq": objId}}

	_, err := col.UpdateOne(ctx, filter, updtString)

	if err != nil {
		return false, err
	}

	return true, nil
}
