package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", indexHandler)
	r.Post("/hook", pullHandler)

	http.ListenAndServe(":3009", r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func pullHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("damn it!"))
}
