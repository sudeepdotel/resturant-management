// main.go
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sudeepdotel/resturant-management/routes"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRestaurantRoutes(router)
	fmt.Println("Listening at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
