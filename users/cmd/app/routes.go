package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/users/", app.getAllUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", app.findUserByID).Methods("GET")
	r.HandleFunc("/api/users/", app.insertUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", app.deleteUser).Methods("DELETE")

	r.HandleFunc("/api/registredUsers/", app.getAllRegistredUsers).Methods("GET")
	r.HandleFunc("/api/registredUsers/{id}", app.findRegistredUserByID).Methods("GET")
	r.HandleFunc("/api/registredUsers/", app.insertRegistredUser).Methods("POST")
	r.HandleFunc("/api/registredUsers/{id}", app.deleteRegistredUser).Methods("DELETE")

	r.HandleFunc("/api/verifications/", app.getAllVerifications).Methods("GET")
	r.HandleFunc("/api/verifications/{id}", app.findVerificationByID).Methods("GET")
	r.HandleFunc("/api/verifications/", app.insertVerification).Methods("POST")
	r.HandleFunc("/api/verifications/{id}", app.deleteVerification).Methods("DELETE")

	r.HandleFunc("/api/reports/", app.getAllReports).Methods("GET")
	r.HandleFunc("/api/reports/{id}", app.findReportByID).Methods("GET")
	r.HandleFunc("/api/reports/", app.insertReports).Methods("POST")
	r.HandleFunc("/api/reports/{id}", app.deleteReport).Methods("DELETE")


	r.HandleFunc("/api/roles/", app.getAllRoles).Methods("GET")
	r.HandleFunc("/api/roles/{id}", app.findRoleByID).Methods("GET")
	r.HandleFunc("/api/roles/", app.insertRole).Methods("POST")
	r.HandleFunc("/api/roles/{id}", app.deleteRole).Methods("DELETE")

	r.HandleFunc("/api/agents/", app.getAllAgents).Methods("GET")
	r.HandleFunc("/api/agents/{id}", app.findAgentByID).Methods("GET")
	r.HandleFunc("/api/agents/", app.insertAgent).Methods("POST")
	r.HandleFunc("/api/agents/{id}", app.deleteAgent).Methods("DELETE")

	return r
}
