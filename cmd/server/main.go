package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rgoncalvesrr/fullcycle-clean-arch/configs"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/dto"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/entity"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg := configs.LoadConfig(".")

	fmt.Println(cfg.JWTSecret)

	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productGateway := database.NewProductGateway(db)
	productHandler := NewProductHandler(productGateway)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.WebServerPort), nil)
}

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
