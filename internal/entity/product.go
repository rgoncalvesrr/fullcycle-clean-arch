package entity

import (
	"errors"
	"time"

	"github.com/rgoncalvesrr/fullcycle-clean-arch/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrInvalidID       = errors.New("invalid id")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	p := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := p.Validate()

	if err != nil {
		p = nil
	}

	return p, err
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price == 0 {
		return ErrPriceIsRequired
	}

	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}
