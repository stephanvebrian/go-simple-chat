package main

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stphanvebrian/go-simple-chat/usecase/handler"
)

func startServer(db *sql.DB) error {
	router := mux.NewRouter()

	handler.NewHome(router)
	handler.NewChat(router, db)

	return http.ListenAndServe(":8080", router)
}
