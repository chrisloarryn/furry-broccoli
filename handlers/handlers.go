package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ccontreras/furry-broccoli/middlew"
	"github.com/ccontreras/furry-broccoli/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers for manage some routes, sets the port and start the server
func Handlers() {
	router := mux.NewRouter()

	apiCreate := router.PathPrefix("/api/v1").Subrouter()

	// Routes for the beers -> api/v1/beers
	apiCreate.HandleFunc("/beers", middlew.CheckDB(routers.SearchBeers)).Methods(http.MethodGet)
	apiCreate.HandleFunc("/beers", middlew.CheckDB(routers.SaveBeer)).Methods(http.MethodPost)
	apiCreate.HandleFunc("/beers/{beerId:[0-9]+}", middlew.CheckDB(routers.SearchBeerByBeerId)).Methods(http.MethodGet)
	apiCreate.HandleFunc("/beers/{beerId:[0-9]+}/boxprice", middlew.CheckDB(routers.BoxBeerPriceByBeerId)).Methods(http.MethodGet)

	// listings -> api/v1/listings
	apiCreate.HandleFunc("/listings", middlew.CheckDB(routers.SearchListingAndReview)).Methods(http.MethodGet)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
