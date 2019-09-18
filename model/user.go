package model

import (
	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
	"top.guitoubing/gotest/db"
)

type User struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

func EncryptPassword(password string) (encrypted string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *User) GenerateID() {
	if u.Id == "" {
		u.Id = bson.NewObjectId()
	}
}

func (u *User) PersisPre()  {}

func (u *User) APIResponsePre() {}

func (u *User) CollectionName() string {
	return db.CUser
}

func (u *User) CryptPassword() error {
	if u.Password != "" {
		return u.SetPassword(u.Password)
	}
	return nil
}

func (u *User) SetPassword(password string) error {
	encrypted, err := EncryptPassword(password)
	if err != nil {
		return err
	}
	u.Password = encrypted
	return nil
}

func (u *User) ComparePassword(password string ) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

