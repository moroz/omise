package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/moroz/omise/config"
	"github.com/moroz/omise/controllers"
)

var db = sqlx.MustConnect("postgres", config.DATABASE_URL)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", controllers.ProductIndex(db))
	r.Get("/products/new", controllers.NewProduct)
	r.Get("/products/{slug}", controllers.ShowProduct(db))
	r.Post("/products", controllers.CreateProduct(db))

	if config.ProductionMode {
		fs := http.FileServer(http.Dir("./static/assets"))
		r.Handle("/assets/*", http.StripPrefix("/assets/", fs))
	}

	fmt.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
