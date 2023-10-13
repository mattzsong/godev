package data

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                string `bson:"_id,omitempty" json:"id"`
	Username          string `bson:"username" json:"username"`
	Email             string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"encryptedPassword" json:"encryptedPassword"`
	IsAdmin           bool   `bson:"isAdmin" json:"isAdmin"`
}

func NewUser(username, email, password string) (*User, error) {
	encrypted_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Username:          username,
		Email:             email,
		EncryptedPassword: string(encrypted_password),
	}, nil
}

func NewAdminUser(username, email, password string) (*User, error) {
	user, err := NewUser(username, email, password)
	if err != nil {
		return nil, err
	}
	user.IsAdmin = true
	return user, nil
}

func (u *User) ValidatePassword(pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pw))
	return err == nil
}
