package database

import (
	"testing"

	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Usuario X", "usuario@dominio.com", "12345")
	userDB := NewUserGateway(db)
	err = userDB.Create(user)

	assert.Nil(t, err)

	var userFound entity.User

	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.EMail, userFound.EMail)
	assert.NotEmpty(t, userFound.Password)
}

func TestFindByEmailUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
		return
	}
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		t.Error(err)
		return
	}
	user, _ := entity.NewUser("Usuario X", "usuario@dominio.com", "12345")
	userDB := NewUserGateway(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.EMail)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.EMail, userFound.EMail)
	assert.NotEmpty(t, userFound.Password)
}
