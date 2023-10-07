package dao

import (
	"context"
	
	"speed_quest_server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type ApplicationAnswersDAO struct {
	Collection *mongo.Collection
	Context context.Context
	
}

func (dao ApplicationAnswersDAO) GetByVisaApplicationID(id primitive.ObjectID) (*models.ApplicationAnswer, error){

	var applicationAnswers models.ApplicationAnswer
	err := dao.Collection.FindOne(dao.Context, bson.M{"_id":id} ).Decode(&applicationAnswers)
	if err != nil {
		println(err.Error())
		return nil, err
	}	

	return &applicationAnswers, nil
}