package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/notifications/", app.getAllNotification).Methods("GET")
	r.HandleFunc("/api/notifications/{id}", app.findByIDNotification).Methods("GET")
	r.HandleFunc("/api/notifications/", app.insertNotification).Methods("POST")
	r.HandleFunc("/api/notifications/{id}", app.deleteNotification).Methods("DELETE")

	return r
}
