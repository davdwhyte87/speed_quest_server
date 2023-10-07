package models

// this helps us know if a user is authenticated 

type Auth struct{
	UserName string `bson:"user_name"`
	UserId string `bson:"user_id"`
	Token string `bson:"token"`
}
