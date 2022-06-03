package main

import (
	mux2 "github.com/gorilla/mux"
	"joiner-management/db"
	"joiner-management/handlers"
	"log"
	"net/http"
)

func main() {

	DB := db.InitDb()
	defaultHandler := handlers.New(DB)

	router := mux2.NewRouter()

	router.HandleFunc("/joiner", defaultHandler.AddJoiner).Methods(http.MethodPost)
	router.HandleFunc("/joiner", defaultHandler.GetJoiners).Methods(http.MethodGet)
	router.HandleFunc("/joiner/{id}", defaultHandler.DeleteJoiner).Methods(http.MethodDelete)
	router.HandleFunc("/joiner/{id}", defaultHandler.UpdateJoiner).Methods(http.MethodPut)
	router.HandleFunc("/joiner/{id}", defaultHandler.GetJoiner).Methods(http.MethodGet)

	log.Println("Joiner server is running")
	http.ListenAndServe(":8000", router)
}
