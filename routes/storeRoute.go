package routes

import (
	"github.com/gorilla/mux"
	"github.com/subhammahanty235/store-api-golang/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/items", controllers.GetAllItems).Methods("GET")
	router.HandleFunc("/api/insert", controllers.InsertOneItem).Methods("POST")

	router.HandleFunc("/api/update/{id}", controllers.UpdateOneItem).Methods("PATCH")

	return router

}
