package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Entry (w http.ResponseWriter, x *http.Request) {
	json.NewEncoder(w).Encode("Welcome Index")
}

func main () {
	router := mux.NewRouter()

	router.HandleFunc("/", Entry).Methods("GET");
	//router.HandleFunc("/users", CreatePost).Methods("POST")
	//router.HandleFunc("/auth/login", CreatePost).Methods("POST")
	
	//router.HandleFunc("/post/create", CreatePost).Methods("POST")

	fmt.Println("Server Running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}