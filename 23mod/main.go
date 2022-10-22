package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Go Mod")
	r := mux.NewRouter()
	r.HandleFunc("/", server).Methods("GET")
	log.Fatal(http.ListenAndServe(":4001", r))
}
func server(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to go"))
}
