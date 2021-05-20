package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RetrieveFollowers struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId         string             `bson:"UserId" json:"userId,omitempty"`
	UserRelationId string             `bson:"UserRelationId" json:"userRelationId,omitempty"`
	Tweet          struct {
		Message string    `bson:"Message" json:"message,omitempty"`
		Date    time.Time `bson:"Date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
