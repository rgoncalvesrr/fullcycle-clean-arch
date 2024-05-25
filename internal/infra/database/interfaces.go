package database

import "github.com/rgoncalvesrr/fullcycle-clean-arch/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
