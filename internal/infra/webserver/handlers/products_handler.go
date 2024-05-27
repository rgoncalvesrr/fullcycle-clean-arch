package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/dto"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/entity"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/infra/database"
)

type ProductHandler struct {
	ProductGateway database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductGateway: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&productDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(productDto.Name, productDto.Price)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.ProductGateway.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	o := &dto.CreateProductOutput{
		ID:        p.ID.String(),
		Name:      p.Name,
		Price:     p.Price,
		CreatedAt: p.CreatedAt,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&o)
}
