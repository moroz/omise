package controllers

import (
	"net/http"

	"github.com/moroz/omise/templates"
)

func ProductIndex(w http.ResponseWriter, r *http.Request) {
	templates.ProductIndex().Render(r.Context(), w)
}
