package dao

import (
	"context"
	"speed_quest_server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (pdao *PlayerStatsDAO) GetLeaderBoard()([]models.PlayerStats, error){
	var playerStats []models.PlayerStats
	opts:= options.Find()
	opts.SetSort(bson.M{"high_score":-1})
	opts.SetLimit(3)
    cursor, err :=pdao.Collection.Find(pdao.Context, bson.M{}, opts)
	err =cursor.All(pdao.Context, &playerStats)
	return playerStats, err
	
}