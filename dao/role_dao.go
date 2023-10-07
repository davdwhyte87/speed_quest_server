package dao

import (
	"context"
	"speed_quest_server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type RoleDAO struct {
	Collection *mongo.Collection
	Context context.Context
}

// get the role for a particular user based on id 
func (roleDAO *RoleDAO) GetUserRole(roleID int) ( models.Role){
	// 
	var role models.Role

	roleDAO.Collection.FindOne(roleDAO.Context, bson.M{"_id":roleID} ).Decode(&role)
	return role
}