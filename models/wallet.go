package models



type Wallet struct {
	Id string `bson:"id"`
	Balance int64 `bson:"balance"`
	UserId string `bson:"user_id"`
	Address string `bson:"address"`
	Email string `bson:"email"`
	UpdatedAt string `bson:"updated_at"`
	CreatedAt string `bson:"created_at"`
	Blocks []Block `bson:"blocks"`
	Version float64 `bson:"version"` // 1.0, 1.1, 1.2, 1.3 ...
}

type Block struct {
	Id string
	SenderAddress string 
	Amount int64
	ReceiverAddress string 
	Date string
	PreviousHash string 
}



