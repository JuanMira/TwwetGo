package models

type Relation struct {
	UserId         string `bson:"UserID" json:"userID"`
	UserRelationId string `bson:"UserRelationId" json:"userRelationId"`
}
