package model

type User struct{
	UserID		string	`json:userid`
	Username	string	`json:username`
	Email		string	`json:email`
	Password	string	`json:password`
}