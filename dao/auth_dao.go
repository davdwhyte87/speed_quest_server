package dao

import (
	"context"
	"speed_quest_server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)



type AuthDAO struct {
	Collection *mongo.Collection
	Context context.Context
}

// get player stats by user id
func (adao *AuthDAO) GetByToken(id string) ( models.Auth, error){
	// 
	var auth models.Auth

	err := adao.Collection.FindOne(adao.Context, bson.M{"token":id} ).Decode(&auth)
	return auth, err
}

func (adao *AuthDAO) DeleteByUserName(userName string) error {
	_, err := adao.Collection.DeleteMany(adao.Context, bson.M{"user_name":userName})
	return err
}

// update player stats
// func  (wdao *WalletDAO) Update(id string, wallet models.Wallet ) error {
// 	_, err := wdao.Collection.UpdateOne(wdao.Context, bson.M{"id": id}, bson.M{"$set": wallet})
// 	return err
// }

