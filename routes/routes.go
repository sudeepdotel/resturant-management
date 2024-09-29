package routes

import (
	"github.com/gorilla/mux"
	"github.com/sudeepdotel/resturant-management/handlers"
)

func RegisterRestaurantRoutes(router *mux.Router) {
	router.HandleFunc("/restaurants", handlers.GetRestaurants).Methods("GET")
	router.HandleFunc("/restaurants", handlers.CreateRestaurants).Methods("POST")
}
