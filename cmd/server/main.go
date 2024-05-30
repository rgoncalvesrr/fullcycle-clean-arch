package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/configs"
	_ "github.com/rgoncalvesrr/fullcycle-clean-arch/docs"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/entity"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/infra/database"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//	@title			Go Expert API Example
//	@version		1.0
//	@description	Product API with authentication
//	@termsOfService	https://swagger.io/terms/

//	@contact.name	Ricardo Gonçalves
//	@contact.url	https://goncalves.biz
//	@contact.email	ricardo@goncalves.biz

// @host						localhost:8080
// @BasePath					/
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	cfg := configs.LoadConfig(".")

	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productGateway := database.NewProductGateway(db)
	productHandler := handlers.NewProductHandler(productGateway)

	userGateway := database.NewUserGateway(db)
	userHandler := handlers.NewUserHandler(userGateway, cfg.TokenAuth, cfg.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use(LogRequest)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth)) // verificação do token JWT
		r.Use(jwtauth.Authenticator)           // validação do token
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Post("/", productHandler.CreateProduct)
		r.Patch("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/auth", userHandler.GetJWT)

	// r.Route("", func(r chi.Router) {
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(fmt.Sprintf("http://localhost:%s/docs/doc.json", cfg.WebServerPort))))
	// })

	http.ListenAndServe(fmt.Sprintf(":%s", cfg.WebServerPort), r)
}

func LogRequest(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, " - ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
