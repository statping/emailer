package main

import "fmt"

func FindUser(email string) (*User, error) {
	var u User
	err := db.Model(&User{}).Where("email = ?", email).Find(&u)
	if err.Error != nil {
		return nil, err.Error
	}
	return &u, nil
}

func FindUserKey(key string) (*User, error) {
	var u User
	err := db.Model(&User{}).Where("key = ?", key).Find(&u)
	if err.Error != nil {
		return nil, err.Error
	}
	return &u, nil
}

func (u *User) Create() error {
	e := db.Model(&User{}).Create(u)
	return e.Error
}

func (u *User) Update() error {
	e := db.Model(&User{}).Update(u)
	return e.Error
}

func (u *User) Delete() error {
	e := db.Model(&User{}).Delete(u)
	return e.Error
}

func (u *User) ConfirmLink() string {
	return fmt.Sprintf("https://emailer.statping.com/confirm/%s", u.Key)
}
