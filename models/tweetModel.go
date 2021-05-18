package models

type TweetModel struct {
	Message string `bson:"Message" json:"message"`
}
