package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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

// Create Product godoc
//
//	@Summay			Create Product
//	@Description	Create Product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//
//	@Param			request	body		dto.CreateProductInput	true	"product request"
//	@Success		201		{object}	dto.CreateProductOutput
//	@Failure		500		{object}	dto.Error
//	@Router			/products [post]
//
//	@Security		ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&productDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&dto.Error{Message: err.Error()})
		return
	}

	p, err := entity.NewProduct(productDto.Name, productDto.Price)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&dto.Error{Message: err.Error()})
		return
	}

	err = h.ProductGateway.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&dto.Error{Message: err.Error()})
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

// Get Product godoc
//
//	@Summay			Get Product
//	@Description	Get Product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path		string	true	"Product ID"	Format(uuid)
//	@Success		200	{object}	entity.Product
//	@Failure		404	{object}	dto.Error
//	@Router			/products/{id} [get]
//
//	@Security		ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductGateway.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&dto.Error{Message: err.Error()})
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

// Update Product godoc
//
//	@Summay			Update Product
//	@Description	Update Product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//
//	@Param			id		path		string					true	"Product ID"	Format(uuid)
//	@Param			request	body		dto.CreateProductInput	true	"Product"
//	@Success		200		{object}	entity.Product
//	@Failure		400		{object}	dto.Error
//	@Failure		404		{object}	dto.Error
//	@Failure		500		{object}	dto.Error
//	@Router			/products/{id} [put]
//
//	@Security		ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductGateway.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&dto.Error{Message: err.Error()})
		return
	}

	var productDto dto.CreateProductInput
	err = json.NewDecoder(r.Body).Decode(&productDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&dto.Error{Message: err.Error()})
		return
	}

	product.Name = productDto.Name
	product.Price = productDto.Price

	err = h.ProductGateway.Update(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&dto.Error{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Delete Product godoc
//
//	@Summay			Delete Product
//	@Description	Delete Product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path	string	true	"Product ID"	Format(uuid)
//	@Success		204
//	@Failure		400
//	@Failure		404	{object}	dto.Error
//	@Router			/products/{id} [delete]
//
//	@Security		ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.ProductGateway.Delete(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&dto.Error{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// List Product godoc
//
//	@Summay			List Products
//	@Description	List Products
//	@Tags			products
//	@Accept			json
//	@Produce		json
//
//	@Param			page	query	string	false	"page number"
//	@Param			limit	query	string	false	"limit"
//	@Param			sort	query	string	false	"order type"
//	@Success		200		{array}	entity.Product
//	@Success		204
//	@Failure		500	{object}	dto.Error
//	@Router			/products [get]
//
//	@Security		ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 50
	}

	offset := (pageInt - 1) * limitInt

	products, err := h.ProductGateway.FindAll(offset, limitInt, sort)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&dto.Error{Message: err.Error()})
		return
	}

	if len(products) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&products)
}
