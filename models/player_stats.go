package models

type PlayerStats struct {
	Id        string `bson:"id"`
	UserID    string `bson:"user_id"`
	UserName  string `bson:"user_name"`
	HighScore int64  `bson:"high_score"`
	LastScore int64  `bson:"last_score"`
	CreateAt  string `bson:"created_at"`
	UpdatedAt string `bson:"updated_at"`
}
