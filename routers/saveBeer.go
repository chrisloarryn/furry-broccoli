package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chrisloarryn/furry-broccoli/bd"
	"github.com/chrisloarryn/furry-broccoli/models"
)

// SaveBeer is for saving url handled
func SaveBeer(w http.ResponseWriter, r *http.Request) {
	var bottle models.Beer
	err := json.NewDecoder(r.Body).Decode(&bottle)
	fmt.Println(bottle)
	if err != nil {
		http.Error(w, "An error has occurred when trying to decode. "+err.Error(), 400)
		return
	}
	register := models.SaveBeer{
		BeerID: bottle.BeerID,
		Name: bottle.Name,
		Brewery: bottle.Brewery,
		Country: bottle.Country,
		Price: bottle.Price,
		Currency: bottle.Currency,
	}
	beer, status, err := bd.InsertBeer(register)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "tweet has not been inserted", 400)
		return
	}


	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(beer)
}
