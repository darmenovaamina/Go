package main

import (
	"Assignment2/Web/apihelper"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/User", apihelper.GetAllUserHandler).Methods("GET")
	router.HandleFunc("/api/User", apihelper.AddUserHandler).Methods("POST")
	router.HandleFunc("/api/User/{id}", apihelper.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/api/User/{id}", apihelper.DeleteUserHandler).Methods("DELETE")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
