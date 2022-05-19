package api

import (
	"context"
	"encoding/json"
	"fmt"
	"inventory/db"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.Background()

type InventoryController struct{}

func handleException(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, fmt.Sprintf("Error: %d %s\n", statusCode, err.Error()))
}

func (controller *InventoryController) CreateItem(w http.ResponseWriter, r *http.Request) {
	item := db.InventoryItem{}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		handleException(w, err, http.StatusBadRequest)
		return
	}
	item.ItemId = uuid.New().String()
	id, err := db.Insert(ctx, item)
	if err != nil {
		handleException(w, err, http.StatusInternalServerError)
		return
	}
	response := struct {
		ItemId string `json:"itemId"`
	}{ItemId: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (controller *InventoryController) EditItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["itemId"]
	editableFields := struct {
		ItemName        *string `json:"itemName,omitempty"`
		ItemDescription *string `json:"itemDescription,omitempty"`
	}{}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&editableFields)
	if err != nil {
		handleException(w, err, http.StatusBadRequest)
		return
	}
	update := bson.M{}
	if editableFields.ItemName != nil {
		update["itemName"] = *editableFields.ItemName
	}
	if editableFields.ItemDescription != nil {
		update["itemDescription"] = *editableFields.ItemDescription
	}
	err = db.Update(ctx, itemId, update)
	if err != nil {
		handleException(w, err, http.StatusInternalServerError)
		return
	}
}

func (controller *InventoryController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["itemId"]
	deleteItemRequest := struct {
		Comment *string `json:"comment,omitempty"`
	}{}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&deleteItemRequest)
	if err != nil {
		handleException(w, err, http.StatusBadRequest)
		return
	}
	update := bson.M{"deleted": true}
	if deleteItemRequest.Comment != nil {
		update["deletionComment"] = *deleteItemRequest.Comment
	}
	err = db.Update(ctx, itemId, update)
	if err != nil {
		handleException(w, err, http.StatusInternalServerError)
		return
	}
}

func (controller *InventoryController) UndeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["itemId"]
	err := db.Update(ctx, itemId, bson.M{"deleted": false})
	if err != nil {
		handleException(w, err, http.StatusInternalServerError)
		return
	}
}

func (controller *InventoryController) ListItems(w http.ResponseWriter, r *http.Request) {
	items, err := db.List(ctx, bson.M{"deleted": false})
	if err != nil {
		handleException(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (controller *InventoryController) ListDeletedItems(w http.ResponseWriter, r *http.Request) {
	items, err := db.List(ctx, bson.M{"deleted": true})
	if err != nil {
		handleException(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
