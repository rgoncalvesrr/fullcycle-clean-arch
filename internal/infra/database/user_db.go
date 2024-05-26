package database

import (
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/entity"
	"gorm.io/gorm"
)

type UserGateway struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *UserGateway {
	return &UserGateway{DB: db}
}

func (u *UserGateway) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *UserGateway) FindByEmail(email string) (*entity.User, error) {
	var user *entity.User

	err := u.DB.Where("e_mail = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
