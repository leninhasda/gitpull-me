package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type payload struct {
	Zen    string `json:"zen"`
	HookID int64  `json:"hook_id"`
	Hook   struct {
		Config struct {
			Secret string `json:"secret"`
			URL    string `json:"url"`
		} `json:"config"`
	} `json:"hook"`
}

func pullHandler(w http.ResponseWriter, r *http.Request) {
	h := r.Header

	body := payload{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err)
	}

	object := map[string]interface{}{
		"header":  h,
		"payload": body,
	}
	b, _ := json.Marshal(object)

	f, err := os.OpenFile(fmt.Sprintf("%d", body.HookID), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	n, _ := f.Write(b)

	res := &response{200, n}
	res.json(w)
	return
}
