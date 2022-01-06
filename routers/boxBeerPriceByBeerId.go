package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chrisloarryn/furry-broccoli/bd"
	"github.com/gorilla/mux"
)

// Me is a function to get the profile of current user
func BoxBeerPriceByBeerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// fmt.Println(vars)

	curr := r.URL.Query().Get("currency")
	var qty int

	if len(curr) < 1 {
		http.Error(w, "should send currency parameter ", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("quantity")) > 1 {
		quantity, err := strconv.Atoi(r.URL.Query().Get("quantity"))
		if err != nil {
			http.Error(w, "should send quantity parameter with number upper than zero", http.StatusBadRequest)
			return
		} else {
			qty = quantity
		}
	} else {
		qty = 6
	}

	strVar := vars["beerId"]
	intVar, err := strconv.Atoi(strVar); if err != nil {
		http.Error(w, "Should send the beerId parameter as integer", http.StatusBadRequest)
		return
	}

	price, err := bd.BoxBeerPriceByBeerId(intVar, curr, qty)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(price)
}
