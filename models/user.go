package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id, omitempty" json:"ID"`
	Name      string             `bson:"Name" json:"Name, omitempty"`
	LastName  string             `bson:"LastName" json:"LastName, omitempty"`
	BirthDate time.Time          `bson:"BirthDate" json:"BirthDate, omitempty"`
	Email     string             `bson:"Email" json:"Email"`
	Password  string             `bson:"Password" json:"Password, omitempty"`
	Avatar    string             `bson:"Avatar" json:"Avatar,omitEmpty"`
	Banner    string             `bson:"Banner" json:"Banner,omitEmpty"`
	Location  string             `bson:"Location" json:"Location,omitEmpty"`
	WebSite   string             `bson:"WebSite" json:"WebSite,omitEmpty"`
}
