package handler

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func NewHome(router *mux.Router) *mux.Router {

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", homeHandler).Methods("GET")

	return router
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Info().Err(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Info().Err(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
