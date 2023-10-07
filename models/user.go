package models

type User struct {
	Id       string `bson:"id" `
	UserName string `bson:"user_name" `
	Email    string `bson:"email `
	Code     string `bson:"code `
	UserRoleId int `bson:"user_role_id"`
}

type UserType int

const (
	Player UserType = iota
	Admin
)

func (user *User)Safe() User{
	user.Code = ""
	return *user
}