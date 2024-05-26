package database

import (
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user *entity.User

	err := u.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
