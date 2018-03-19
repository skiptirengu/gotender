package models

import (
	"github.com/skiptirengu/gotender/database"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint    `json:"id"       gorm:"column:user_id;primary_key"`
	Username string  `json:"username" gorm:"column:username"`
	Email    string  `json:"email"    gorm:"column:email"`
	Password string  `json:"-"        gorm:"column:password"`
	Tokens   []Token `json:"tokens"   gorm:"ForeignKey:user_id;AssociationForeignKey:user_id"`
	Admin    bool    `json:"-"        gorm:"column:admin;default:false"`
}

func (u *User) Create() (*gorm.DB) {
	return database.Open().Create(u)
}

func (u *User) SetPassword(pwd string) (error) {
	if bytePwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		u.Password = string(bytePwd)
		return nil
	}
}

func (u *User) ValidatePassword(pwd string) (error) {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
}

func FindUserByEmailOrUsername(email string, username string) (*User) {
	user := &User{}
	if database.Open().Where("email = ? OR username = ?", email, username).First(user).RecordNotFound() {
		return nil
	} else {
		return user
	}
}
