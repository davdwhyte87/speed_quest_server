package models



type Wallet struct {
	Id string `bson:"id"`
	Balance int64 `bson:"balance"`
	UserId string `bson:"user_id"`
	UpdatedAt string `bson:"updated_at"`
	CreatedAt string `bson:"created_at"`
	
}