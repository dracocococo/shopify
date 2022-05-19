package api

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Route() {
	r := mux.NewRouter()
	inventory_controller := InventoryController{}
	r.HandleFunc("/api/inventory/", inventory_controller.CreateItem).Methods("POST")
	r.HandleFunc("/api/inventory/", inventory_controller.ListItems).Methods("GET")
	r.HandleFunc("/api/inventory/{itemId}", inventory_controller.EditItem).Methods("PATCH")
	r.HandleFunc("/api/inventory/{itemId}", inventory_controller.DeleteItem).Methods("DELETE")
	r.HandleFunc("/api/inventory/recover/", inventory_controller.ListDeletedItems).Methods("GET")
	r.HandleFunc("/api/inventory/recover/{itemId}", inventory_controller.UndeleteItem).Methods("POST")

	// for demonstration purpose just allow access from anywhere
	validator := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PATCH", "DELETE", "OPTIONS"})
	http.ListenAndServe(":9876", handlers.CORS(validator, headers, methods)(r))
}
