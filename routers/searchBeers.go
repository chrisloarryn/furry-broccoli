package routers

import (
	"encoding/json"
	"net/http"

	"github.com/chrisloarryn/furry-broccoli/bd"
)

// SearchBeers for reading tweets
func SearchBeers(w http.ResponseWriter, r *http.Request) {
	response, ok := bd.SearchBeers()
	if !ok {
		http.Error(w, "error reading beers", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}
