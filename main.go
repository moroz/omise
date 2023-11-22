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

func OverrideRequestMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			next.ServeHTTP(w, r)
			return
		}
		err := r.ParseForm()
		if err != nil || r.Form["_method"] == nil {
			next.ServeHTTP(w, r)
			return
		}
		switch r.FormValue("_method") {
		case "PUT", "PATCH", "DELETE":
			r.Method = r.FormValue("_method")
		default:
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := chi.NewRouter()
	r.Use(OverrideRequestMethod)
	r.Use(middleware.Logger)
	r.Get("/", controllers.ProductIndex(db))
	r.Get("/products/new", controllers.NewProduct)
	r.Get("/products/{slug}", controllers.ShowProduct(db))
	r.Get("/products/{slug}/edit", controllers.EditProduct(db))
	r.Post("/products", controllers.CreateProduct(db))
	r.Patch("/products/{id}", controllers.UpdateProduct(db))

	if config.ProductionMode {
		fs := http.FileServer(http.Dir("./static/assets"))
		r.Handle("/assets/*", http.StripPrefix("/assets/", fs))
	}

	fmt.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
