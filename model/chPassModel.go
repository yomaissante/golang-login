package model

type ChangePassword struct {
	Username	string	`json:"username"`
	OldPassword	string	`json:"old_password"`
	NewPassword	string	`json:"new_password"`
	ConfirmPassword	string	`json:"confirm_password"`
}