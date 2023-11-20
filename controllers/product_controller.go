package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/moroz/omise/models"
	"github.com/moroz/omise/templates"
)

func ProductIndex(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := models.ListProducts(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Internal Server Error")
			return
		}
		templates.ProductIndex(products).Render(r.Context(), w)
	}
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	templates.NewProduct().Render(r.Context(), w)
}

func CreateProduct(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Bad request")
			return
		}

		name := r.PostForm.Get("name")
		priceStr := r.PostForm.Get("price")
		price, err := strconv.Atoi(priceStr)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Bad request")
			return
		}
		description := r.PostForm.Get("description")
		_, err = models.CreateProduct(db, name, price, description)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintf(w, "An error occurred: %s", err)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func ShowProduct(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")
		product, err := models.GetProductBySlug(db, slug)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			tpl := templates.NotFound()
			tpl.Render(r.Context(), w)
			return
		}

		tpl := templates.ShowProduct(product)
		tpl.Render(r.Context(), w)
	}
}
