package dao

import (
	"context"
	"speed_quest_server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)



type WalletDAO struct {
	Collection *mongo.Collection
	Context context.Context
}

// get player stats by user id
func (wdao *WalletDAO) GetByUserId(id string) ( models.Wallet, error){
	// 
	var wallet models.Wallet

	err :=wdao.Collection.FindOne(wdao.Context, bson.M{"user_id":id} ).Decode(&wallet)
	return wallet, err
}

// update player stats
func  (wdao *WalletDAO) Update(id string, wallet models.Wallet ) error {
	_, err := wdao.Collection.UpdateOne(wdao.Context, bson.M{"id": id}, bson.M{"$set": wallet})
	return err
}