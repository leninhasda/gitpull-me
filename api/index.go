package api

import "net/http"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	res := &response{200, "hello"}
	res.string(w)
	return
}
