package main

import (
	"fmt"
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
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", r)
}
