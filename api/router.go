package api

import (
	"github.com/go-chi/chi"
)

// Router returns a chi router mux
func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", indexHandler)
	r.Post("/hook", pullHandler)
	return r
}
