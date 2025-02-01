package handlers

import (
	"net/http"

	"github.com/zyonson/go-course/pkg/render"
)

/*
Home is the home page handler

	w http.ResponseWriter: This parameter is used to write the HTTP response back to the client.
	r *http.Request: This parameter contains all the information about the HTTP request, such as headers, URL, method, and more.
*/
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
