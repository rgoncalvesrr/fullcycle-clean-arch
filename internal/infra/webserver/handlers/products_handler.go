package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&o)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductGateway.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	o := &dto.CreateProductOutput{
		ID:        product.ID.String(),
		Name:      product.Name,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&o)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductGateway.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var productDto dto.CreateProductInput
	err = json.NewDecoder(r.Body).Decode(&productDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.Name = productDto.Name
	product.Price = productDto.Price

	err = h.ProductGateway.Update(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.ProductGateway.Delete(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
