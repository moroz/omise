package controllers

import (
	"fmt"
	"net/http"

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
