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

// CreateRestaurants CreateRestaurant Add a new restaurant with error handling
func CreateRestaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newRestaurants []models.Restaurant
	if err := json.NewDecoder(r.Body).Decode(&newRestaurants); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Add each new restaurant to the restaurants list
	for _, restaurant := range newRestaurants {
		restaurant.ID = uint(len(restaurants) + 1)
		restaurants = append(restaurants, restaurant)
	}

	// Respond with the newly added restaurants
	if err := json.NewEncoder(w).Encode(newRestaurants); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to add restaurants", http.StatusInternalServerError)
		return
	}
}
