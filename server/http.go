package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"mysql/database"
	"mysql/usage"
	"net/http"
	"os"
)

func Run() {
	port := os.Getenv("APP_PORT")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/categories", categories)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func categories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(database.GetAllCategories())
	usage.CheckErr(err)
}
