package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type User struct {
	Username string
	Password string
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("Required Name")
	}
	if u.Password == "" {
		return errors.New("Required Type")
	}
	return nil

}

func (u *User) Prepare() {
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
}

func SaveUser(db *gorm.DB) (*User, error) {

	var err error
	u := User{"admin", "admin"}
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return &u, nil
}
