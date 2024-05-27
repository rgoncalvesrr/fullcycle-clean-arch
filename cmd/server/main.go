package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/configs"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/entity"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/infra/database"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg := configs.LoadConfig(".")

	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productGateway := database.NewProductGateway(db)
	productHandler := handlers.NewProductHandler(productGateway)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/products/{id}", productHandler.GetProduct)
	r.Post("/products", productHandler.CreateProduct)
	r.Patch("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	http.ListenAndServe(fmt.Sprintf(":%s", cfg.WebServerPort), r)
}
