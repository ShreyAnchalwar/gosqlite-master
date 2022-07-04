package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *App) registerroutes() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/list", app.GetData)
	r.Post("/add", app.AddData)

	return r

}
