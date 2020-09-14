package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/randomUser/apis/userAPI"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/user/findall", userAPI.FindAll).Methods("GET")
	router.HandleFunc("/api/user/search/{keyword}", userAPI.Search).Methods("GET")
	router.HandleFunc("/api/user/add", userAPI.Add).Methods("POST")
	router.HandleFunc("/api/user/remove", userAPI.Remove).Methods("POST")

	err := http.ListenAndServe(":500", router)
	if err != nil {
		fmt.Println(err)
	}
}
