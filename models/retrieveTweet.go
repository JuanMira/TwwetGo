package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RetrieveTweet struct {
	Id      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId  string             `bson:"UserId" json:"userId,omitempty"`
	Message string             `bson:"Message" json:"message,omitempty"`
	Date    time.Time          `bson:"Date" json:"date,omitempty"`
}
