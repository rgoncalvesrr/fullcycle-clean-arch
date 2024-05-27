package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestProductCreate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Produto", 50.0)

	assert.Nil(t, err)
	assert.NotNil(t, product)

	productGateway := NewProductGateway(db)

	assert.NotNil(t, productGateway)
	err = productGateway.Create(product)

	assert.NoError(t, err)
}

func TestProductFindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	productGateway := NewProductGateway(db)
	assert.NotNil(t, productGateway)

	var e error
	var product *entity.Product

	for i := 0; i < 100; i++ {
		product, e = entity.NewProduct(fmt.Sprintf("Produto %d", i+1), rand.Float64()*100)
		assert.NoError(t, e)

		e = productGateway.Create(product)
		assert.NoError(t, e)
	}

	products, err := productGateway.FindAll(0, 100, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 50)
	assert.Equal(t, "Produto 1", products[0].Name)
	assert.Equal(t, "Produto 40", products[39].Name)
	assert.Equal(t, "Produto 50", products[49].Name)

	products, err = productGateway.FindAll(50, 100, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 50)
	assert.Equal(t, "Produto 51", products[0].Name)
	assert.Equal(t, "Produto 90", products[39].Name)
	assert.Equal(t, "Produto 100", products[49].Name)
}

func TestProductFindById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	productGateway := NewProductGateway(db)
	assert.NotNil(t, productGateway)

	var e error
	var product *entity.Product

	for i := 0; i < 10; i++ {
		product, e = entity.NewProduct(fmt.Sprintf("Produto %d", i+1), rand.Float64()*100)
		assert.NoError(t, e)

		e = productGateway.Create(product)
		assert.NoError(t, e)
	}

	products, err := productGateway.FindAll(0, 0, "x")
	assert.NoError(t, err)
	product = &products[0]

	productFound, err := productGateway.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestProductUpdate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	productGateway := NewProductGateway(db)
	assert.NotNil(t, productGateway)

	expectedName := "Produto 1"
	product, err := entity.NewProduct(expectedName, rand.Float64()*100)
	assert.NoError(t, err)

	err = productGateway.Create(product)
	assert.NoError(t, err)

	product, err = productGateway.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, expectedName, product.Name)

	expectedName = "Produto Atualizado"
	product.Name = expectedName

	err = productGateway.Update(product)
	assert.NoError(t, err)
	assert.Equal(t, expectedName, product.Name)

}

func TestProductDelete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	productGateway := NewProductGateway(db)
	assert.NotNil(t, productGateway)

	product, err := entity.NewProduct("Produto", rand.Float64()*100)
	assert.NoError(t, err)

	err = productGateway.Create(product)
	assert.NoError(t, err)

	err = productGateway.Delete(product.ID.String())
	assert.NoError(t, err)
}
