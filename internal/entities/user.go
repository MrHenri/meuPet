package entities

import (
	"errors"

	entityPkg "github.com/MrHenri/meuPet/pkg/entities"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"passoword"`
	Phone    string `json:"phone"`
}

func NewUser(email, password, phone string) (*User, error) {
	id := entityPkg.NewID().String()
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{ID: id, Email: email, Password: string(hash), Phone: phone}
	return user, user.ValidateUserCreation()
}

func (u *User) Authentication(passoword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passoword))
	return err == nil
}

func (u *User) ValidateUserCreation() error {
	if u.ID == "" {
		return errors.New("invalid ID Generated")
	}
	if u.Email == "" {
		return errors.New("invalid Email")
	}
	if u.Password == "" {
		return errors.New("invalid Password")
	}
	return nil
}
