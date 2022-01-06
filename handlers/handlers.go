package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/chrisloarryn/furry-broccoli/middleware"
	"github.com/chrisloarryn/furry-broccoli/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers for manage some routes, sets the port and start the server
func Handlers() {
	router := mux.NewRouter()

	apiCreate := router.PathPrefix("/api/v1").Subrouter()
	
	// Routes for the beers -> api/v1/beers
	apiCreate.HandleFunc("/beers", middleware.CheckDB(routers.SearchBeers)).Methods(http.MethodGet)
	apiCreate.HandleFunc("/beers", middleware.CheckDB(routers.SaveBeer)).Methods(http.MethodPost)
	apiCreate.HandleFunc("/beers/{beerId:[0-9]+}", middleware.CheckDB(routers.SearchBeerByBeerId)).Methods("GET")
	apiCreate.HandleFunc("/beers/{beerId:[0-9]+}/boxprice", middleware.CheckDB(routers.BoxBeerPriceByBeerId)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
