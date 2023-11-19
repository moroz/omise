package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/moroz/omise/controllers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", controllers.ProductIndex)
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", r)
}
