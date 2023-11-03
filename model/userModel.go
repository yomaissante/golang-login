package model

type Userdata struct{
	UserID		string	`json:"userid"`
	Username	string	`json:"username"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}