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

	log.Println("Joiner server is running")
	http.ListenAndServe(":8000", router)
}
