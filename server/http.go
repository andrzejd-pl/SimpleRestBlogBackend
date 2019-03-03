package server

import (
	"crypto/tls"
	"encoding/json"
	"github.com/andrzejd-pl/SimpleRestBlogBackend/database"
	"github.com/andrzejd-pl/SimpleRestBlogBackend/usage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Run() {
	port := os.Getenv("APP_PORT")
	certificate := os.Getenv("SERVER_CRT")
	key := os.Getenv("SERVER_KEY")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/categories", categories)
	router.HandleFunc("/articles", articles)

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	log.Fatal(srv.ListenAndServeTLS(certificate, key))
}

func categories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(database.GetAllCategories())
	usage.CheckErr(err)
}

func articles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(database.GetAllArticles())
	usage.CheckErr(err)
}
