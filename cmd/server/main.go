package main

import (
	"fmt"
	"net/http"

	"github.com/rgoncalvesrr/fullcycle-clean-arch/configs"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/entity"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/infra/database"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/infra/webserver/handlers"
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
	productHandler := handlers.NewProductHandler(productGateway)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.WebServerPort), nil)
}
