package entity

import (
	"github.com/rgoncalvesrr/fullcycle-clean-arch/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	EMail    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name string, email string, password string) (*User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 1)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		EMail:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
