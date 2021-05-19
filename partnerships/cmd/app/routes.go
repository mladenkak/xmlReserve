package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/partnerships/", app.getAllPartnerships).Methods("GET")
	r.HandleFunc("/api/partnerships/{id}", app.findPartnershipByID).Methods("GET")
	r.HandleFunc("/api/bookinpartnershipsgs/", app.insertPartnership).Methods("POST")
	r.HandleFunc("/api/partnerships/{id}", app.deletePartnership).Methods("DELETE")

	return r
}
