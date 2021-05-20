package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/JuanMira/tweetgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var include bool

func ReadUsersAll(Id string, page int64, search string, _type string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tweetGo")
	col := db.Collection("Users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"Name": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var finded bool

	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		var r models.Relation
		r.UserId = Id
		r.UserRelationId = s.ID.Hex()

		include = false
		finded, err = FindRelations(r)

		if _type == "new" && finded == false {
			include = true
		}

		if _type == "follow" && finded == true {
			include = true
		}

		if r.UserRelationId == Id {
			include = false
		}

		if include == true {
			s.Password = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""
			s.WebSite = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cur.Close(ctx)

	return results, true
}
