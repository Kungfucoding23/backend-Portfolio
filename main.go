package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Kungfucoding23/portafolioBackend/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/about", routers.AboutSec)
	router.HandleFunc("/portfolio", routers.PortfolioSection)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
