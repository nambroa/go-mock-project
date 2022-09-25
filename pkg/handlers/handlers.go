package handlers

import (
	"github.com/nambroa/go-mock-project/pkg/render"
	"net/http"
)

// Home is the home page handler.
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml")
}

// About is the about page handler.
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml")
}
