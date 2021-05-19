package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/messages/", app.getAllMessages).Methods("GET")
	r.HandleFunc("/api/messages/{id}", app.findMessageByID).Methods("GET")
	r.HandleFunc("/api/messages/", app.insertMessage).Methods("POST")
	r.HandleFunc("/api/messages/{id}", app.deleteMessage).Methods("DELETE")

	r.HandleFunc("/api/chats/", app.getAllChats).Methods("GET")
	r.HandleFunc("/api/chats/{id}", app.findChatByID).Methods("GET")
	r.HandleFunc("/api/chats/", app.insertChat).Methods("POST")
	r.HandleFunc("/api/chats/{id}", app.deleteChat).Methods("DELETE")

	r.HandleFunc("/api/disposableImages/", app.getAllDisposableImages).Methods("GET")
	r.HandleFunc("/api/disposableImages/{id}", app.findDisposableImageByID).Methods("GET")
	r.HandleFunc("/api/disposableImages/", app.insertDisposableImage).Methods("POST")
	r.HandleFunc("/api/disposableImages/{id}", app.deleteDisposableImage).Methods("DELETE")

	return r
}
