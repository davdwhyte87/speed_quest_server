package dao

import (
	"context"
	"fmt"
	"speed_quest_server/configs"
	"speed_quest_server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type VisaApplicationDAO struct {
	Collection *mongo.Collection
	Context context.Context
	
}
func (dao *VisaApplicationDAO )UpdateVisaApplication(visaApplication models.VisaApplication) {
	// Update an existing user
	// docID, _ := primit.ObjectIDFromHex(user.ID.Hex())
	// _, err := dao.Collection.UpdateOne(dao.ctx, bson.M{"_id": docID}, bson.M{"$set": user})
	// return err
}

// insert 


// get all visa applications 
func (dao *VisaApplicationDAO) GetAll () (interface{}, error){
	pipe := bson.M{
		"$lookup":bson.M{
			"from":"VisaApplicationAnswers",
			"foreignField": "application_id",
			"localField":"_id",
			"as": "application_answers",
		},
	}
	// unwindStage := bson.D{{
	// 	"$lookup", 
	// 	bson.D{
	// 		{"from", "VisaApplications"},
	// 		{"localField",""}
	// 	}}}
	
	pipeline := []bson.M{pipe}
	var visaApplications []bson.M
	cursor, err := dao.Collection.Aggregate(dao.Context, pipeline )
	if err != nil {
		println(err.Error())
		return nil, err
	}
	err = cursor.All(dao.Context, &visaApplications)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	fmt.Printf("%v",visaApplications[0]["application_answers"])

	return visaApplications, nil
}

func (dao *VisaApplicationDAO) GetSingleX(id primitive.ObjectID)(interface{}, error){
	match := bson.M{
		"$match":bson.M{"_id": id },
	}
	pipe := bson.M{
		"$lookup":bson.M{
			"from":"VisaApplicationAnswers",
			"foreignField": "application_id",
			"localField":"_id",
			"as": "application_answers",
		},
	}
	
	pipeline := []bson.M{match,pipe}
	var visaApplications []bson.M
	cursor, err := dao.Collection.Aggregate(dao.Context, pipeline )
	if err != nil {
		println(err.Error())
		return nil, err
	}
	err = cursor.All(dao.Context, &visaApplications)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	fmt.Printf("%v",visaApplications[0]["application_answers"])

	return visaApplications, nil
}

func (dao *VisaApplicationDAO) GetSingle(id primitive.ObjectID)(*models.VisaApplication, error){

	var visaApplication models.VisaApplication
	err := dao.Collection.FindOne(dao.Context, bson.M{"_id":id} ).Decode(&visaApplication)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	// get application answers
	var applicationAnswers []models.ApplicationAnswer
	applicationAnswersCollection := configs.GetCollection(configs.DB, models.ApplicationAnswerCollection)
	cursor, err :=applicationAnswersCollection.Find(dao.Context, bson.M{"application_id": visaApplication.Id})
	if err != nil {
		return nil, err
	}

	err =cursor.All(dao.Context, &applicationAnswers )
	if err != nil {
		return nil, err
	}
	// assign answers data
	visaApplication.VisaApplicationAnswers = applicationAnswers

	// for each application answer get question data
	for index, value := range applicationAnswers{
		// get question
		
		var question models.ApplicationQuestion
		applicationQuestionCollection := configs.GetCollection(configs.DB, models.ApplicationQuestionCollection )
		applicationQuestionCollection.FindOne(dao.Context, bson.M{"_id":value.QuestionId}).Decode(&question)
		value.Question = question
		applicationAnswers[index] = value
	}
	return &visaApplication, nil
}