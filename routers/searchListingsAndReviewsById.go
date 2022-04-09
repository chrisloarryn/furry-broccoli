package routers

import (
	"encoding/json"
	"github.com/ccontreras/furry-broccoli/bd"
	"net/http"
)

// SearchListingAndReview for reading tweets
func SearchListingAndReview(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	//vars := mux.Vars(r)

	if len(ID) < 1 {
		http.Error(w, "Should send the id parameter", http.StatusBadRequest)
		return
	}

	lyr, err := bd.FindListingAndReviewId(ID)
	if err != nil {
		http.Error(w, "An error has occurred when trying to find the profile "+err.Error(), http.StatusBadRequest) //  http.StatusNoContent
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lyr)
}
