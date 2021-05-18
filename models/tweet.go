package models

import "time"

type Tweet struct {
	UserId  string    `bson:"UserId" json:"userId,omitempty"`
	Message string    `bson:"Message" json:"message,omitempty"`
	Date    time.Time `bson:"Date" json:"date,omitempty"`
}
