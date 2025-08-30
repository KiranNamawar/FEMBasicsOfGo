package api

import (
	"encoding/json"
	"museum/data"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["id"]
	if id != nil {
		intId, err := strconv.Atoi(id[0])
		if err == nil && intId < len(data.GetAll()) {
			json.NewEncoder(w).Encode(data.GetAll()[intId])
		} else {
			http.Error(w, "Invalid id", http.StatusBadGateway)
		}
	} else {
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
