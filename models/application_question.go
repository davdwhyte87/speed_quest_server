package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ApplicationQuestion struct {
	Id        primitive.ObjectID `bson:"_id"  json:"id,omitempty"`
	Text      string             `bson:"text"`
	Type      QuestionType       `bson:"type"`
	CreatedAt string             `bson:"created_at"`
}

type QuestionType int

const (
	YesNo QuestionType = iota
	Text
	MultiChoice
)
