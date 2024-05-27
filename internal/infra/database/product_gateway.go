package database

import (
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/entity"
	"gorm.io/gorm"
)

type ProductGateway struct {
	DB *gorm.DB
}

func NewProductGateway(db *gorm.DB) *ProductGateway {
	return &ProductGateway{DB: db}
}

func (p *ProductGateway) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *ProductGateway) FindAll(offset, limit int, sort string) ([]entity.Product, error) {
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if offset < 0 {
		offset = 0
	}

	if limit <= 0 || limit > 50 {
		limit = 50
	}

	var products []entity.Product

	err := p.DB.Limit(limit).Offset(offset).Order("created_at " + sort).Find(&products).Error

	if err != nil {
		products = nil
	}

	return products, err
}

func (p *ProductGateway) FindByID(id string) (*entity.Product, error) {
	var product *entity.Product

	err := p.DB.First(&product, "id = ?", id).Error

	if err != nil {
		product = nil
	}

	return product, err
}

func (p *ProductGateway) Update(product *entity.Product) error {
	if _, err := p.FindByID(product.ID.String()); err != nil {
		return err
	}

	return p.DB.Save(product).Error
}

func (p *ProductGateway) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(&product).Error
}
