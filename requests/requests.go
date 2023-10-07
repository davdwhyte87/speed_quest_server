package requests


type CreateUserReq struct {
	UserName string `validate:"required"`
	Email string `validate:"required,email"`
}

type LoginReq struct{
	Email string `validate:"required,email"`
	Code string 
}

type GetCodeReq struct {
	Email string `validate:"required,email"`
}

type UpdatePlayerStatsReq struct{
	HighScore int64
	LastScore int64
}

type UpdateWalletReq struct{
	Amount int64
}



