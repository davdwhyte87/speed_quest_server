package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type VisaApplication struct {
	Id         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name       string             `json:"name,omitempty" validate:"required"`
	Phone      string             `json:"phone,omitempty" validate:"required"`
	Email      string             `json:"email,omitempty" validate:"required"`
	Location   string             `json:"location,omitempty" validate:"required"`
	Profession string             `json:"profession,omitempty" validate:"required"`
	Status     ApplicationStatus  `bson:"application_status"`
	VisaApplicationAnswers []ApplicationAnswer `bson:"visa_application_answers"`
}

type ApplicationStatus int

const (
	Pending ApplicationStatus = iota
	InReview
	Accepted
	Rejected
)
