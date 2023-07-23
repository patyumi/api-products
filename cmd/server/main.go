package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/patyumi/api-products/configs"
	_ "github.com/patyumi/api-products/docs"
	"github.com/patyumi/api-products/internal/entity"
	"github.com/patyumi/api-products/internal/infra/database"
	"github.com/patyumi/api-products/internal/infra/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DOCUMENTAÇÃO

// @title Go Expert Api Products
// @version 1.0
// @description Product API with authentication
// @termsOfService http://swagger.io/terms/

// @contact.name Patrícia Yumi
// @contact.url https://www.linkedin.com/in/patricia-yumi/
// @contact.email patrciayumi@gmail.com

// @license.name Projeto Pessoal
// @license.url --

// @host localhost:8000
// @BasePath /

// Aqui é porque a gente usa um token gerado para autenticar na API
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JWTExpiresIn))
	//r.Use(LogRequest)

	// agrupamento de rotas
	r.Route("/products", func(r chi.Router) {
		// Middleware que usa o token como método de autenticação
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}

// Middleware exemplo
// Ele recebe a requisição do usuário, faz algo no meio do caminho (nesse caso é o logger) e depois retorna a mesma requisição
// func LogRequest(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Printf("Request: %s %s", r.Method, r.URL.Path)
//		next.ServeHTTP(w, r)
//	})
//}
