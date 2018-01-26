package api

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type response struct {
	StatusCode int
	Data       interface{}
}

func (r *response) string(w http.ResponseWriter) {
	w.WriteHeader(r.StatusCode)
	w.Write([]byte(r.Data.(string)))
	return
}

func (r *response) json(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	if err := json.NewEncoder(w).Encode(r.Data); err != nil {
		panic(err)
	}
	return
}

func (r *response) xml(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/xml")
	if err := xml.NewEncoder(w).Encode(r.Data); err != nil {
		panic(err)
	}
	return
}
