package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/zyonson/go-course/pkg/config"
	"github.com/zyonson/go-course/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
