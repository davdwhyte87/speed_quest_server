package dao

import (
	"context"
	"speed_quest_server/models"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserDAO struct {
	Collection *mongo.Collection
	Context    context.Context
}

// get the role for a particular user based on id
func (userDAO *UserDAO) GetUserRole(roleID int) models.Role {
	//
	var role models.Role

	userDAO.Collection.FindOne(userDAO.Context, bson.M{"_id": roleID}).Decode(&role)
	return role
}

// get user by email

func (userDAO *UserDAO) GetByEmail(email string) (models.User, error) {
	var user models.User
	error := userDAO.Collection.FindOne(userDAO.Context, bson.M{"email": email}).Decode(&user)
	return user, error
}

func (userDAO *UserDAO) GetByUserName(username string) (models.User, error) {
	var user models.User
	error := userDAO.Collection.FindOne(userDAO.Context, bson.M{"user_name": username}).Decode(&user)
	return user, error
}

func (userDAO *UserDAO) GetByUserId(id string) (models.User, error) {
	var user models.User
	error := userDAO.Collection.FindOne(userDAO.Context, bson.M{"id": id}).Decode(&user)
	return user, error
}

func (userDAO *UserDAO) DoesUserExist(email string, username string) (bool, error) {
	var user1 models.User
	var userb models.User
	userDAO.Collection.FindOne(userDAO.Context, bson.M{"email": email}).Decode(&user1)
	userDAO.Collection.FindOne(userDAO.Context, bson.M{"user_name": username}).Decode(&userb)

	// utils.Logger.Debug().Msg(user1.Email)
	if user1.Email == "" && userb.UserName == "" {
		return false, nil
	} else {
		return true, nil
	}
}

// update user model by email
func (userDAO *UserDAO) UpdateByEmail(email string, user models.User) error {

	_, err := userDAO.Collection.UpdateOne(userDAO.Context, bson.M{"email": email}, bson.M{"$set": user})
	return err
}

// insert

// update

// get

// get by id

// delete
