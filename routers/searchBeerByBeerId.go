package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ccontreras/furry-broccoli/bd"
	"github.com/gorilla/mux"
)

// Me is a function to get the profile of current user
func SearchBeerByBeerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strVar := vars["beerId"]
	intVar, err := strconv.Atoi(strVar); if err != nil {
		http.Error(w, "Should send the beerId parameter as integer", http.StatusBadRequest)
		return
	}
	beer, err := bd.FindABeer(intVar)
	if err != nil {
		http.Error(w, "An error has occurred when trying to find the profile "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(beer)
}
