package handlers

import (
	"encoding/json"
	_ "github.com/gorilla/mux"
	"github.com/sudeepdotel/resturant-management/models"
	"log"
	"net/http"
)

var restaurants []models.Restaurant

// GetRestaurants Get all restaurants with error handling
func GetRestaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(restaurants); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to retrieve restaurants", http.StatusInternalServerError)
		return
	}
}

// CreateRestaurant Add a new restaurant with error handling
func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var restaurant models.Restaurant
	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	restaurant.ID = uint(len(restaurants) + 1)
	restaurants = append(restaurants, restaurant)

	if err := json.NewEncoder(w).Encode(restaurant); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to add restaurant", http.StatusInternalServerError)
		return
	}
}
