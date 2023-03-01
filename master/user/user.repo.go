package user

import (
	"gmail-clone.wisnu.net/database"
	"gmail-clone.wisnu.net/modules"
);

type User modules.User

func (user *User) CreateUser(u *User) error {
	result := database.DB.Create(&u)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (user *User) CheckUserByEmail(email string) (*User){
	data := User{}
	result := database.DB.Where(&User{Email: email}).Find(&data)
	if result.RowsAffected == 0 {
		return nil
	}

	return &data
}

func (user *User) CheckUserByPassword(password string) (*User){
	data := User{}
	result := database.DB.Where(&User{Password: password}).Find(&data)
	if result.RowsAffected == 0 {
		return nil
	}

	return &data
}