package routers

import (
	"encoding/json"
	"fmt"
	"github.com/ccontreras/furry-broccoli/bd"
	"github.com/gorilla/mux"
	"net/http"
)

// SearchListingAndReview for reading tweets
func SearchListingAndReview(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	vars := mux.Vars(r)
	fmt.Println(vars)
	fmt.Println(vars)

	if len(ID) < 1 {
		http.Error(w, "Should send the id parameter", http.StatusBadRequest)
		return
	}

	lyr, err := bd.FindListingAndReviewId(ID)
	if err != nil {
		http.Error(w, "An error has occurred when trying to find the profile "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(lyr)
}
