package dao

import (
	"context"
	"speed_quest_server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)



type PlayerStatsDAO struct {
	Collection *mongo.Collection
	Context context.Context
}

// get player stats by user id
func (pdAO *PlayerStatsDAO) GetByUserId(id string) ( models.PlayerStats, error){
	// 
	var playerstats models.PlayerStats

	err :=pdAO.Collection.FindOne(pdAO.Context, bson.M{"user_id":id} ).Decode(&playerstats)
	return playerstats, err
}

// update player stats
func  (pDAO *PlayerStatsDAO) Update(id string, playerStats models.PlayerStats ) error {
	_, err := pDAO.Collection.UpdateOne(pDAO.Context, bson.M{"id": id}, bson.M{"$set": playerStats})
	return err
}