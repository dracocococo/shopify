package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const INVENTORY_COLLECTION = "inventory"

var once sync.Once
var client *mongo.Client
var database *mongo.Database

type InventoryItem struct {
	ItemId          string `json:"itemId" bson:"_id"`
	ItemName        string `json:"itemName" bson:"itemName"`
	ItemDescription string `json:"itemDescription" bson:"itemDescription"`
	Deleted         bool   `json:"deleted" bson:"deleted"`
	DeletionComment string `json:"deletionComment" bson:"deletionComment"`
}

func getDB() (*mongo.Client, *mongo.Database) {
	if client != nil && database != nil {
		return client, database
	}
	once.Do(func() {
		ctx := context.TODO()
		URI := "mongodb://localhost:27017"
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
		if err != nil {
			panic(err)
		}

		if err = client.Ping(ctx, nil); err != nil {
			panic(err)
		}
		database = client.Database("app")
	})
	return client, database
}

func Update(ctx context.Context, itemId string, update bson.M) error {
	_, db := getDB()
	return db.Collection(INVENTORY_COLLECTION).FindOneAndUpdate(ctx, bson.M{"_id": itemId}, bson.M{"$set": update}).Err()
}

func Insert(ctx context.Context, item InventoryItem) (string, error) {
	_, db := getDB()
	res, err := db.Collection(INVENTORY_COLLECTION).InsertOne(ctx, item)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(string), err
}

func Delete(ctx context.Context, itemId string) error {
	_, db := getDB()
	_, err := db.Collection(INVENTORY_COLLECTION).DeleteOne(ctx, bson.M{"_id": itemId})
	return err
}

func Get(ctx context.Context, itemId string) (InventoryItem, error) {
	_, db := getDB()
	item := InventoryItem{}
	err := db.Collection(INVENTORY_COLLECTION).FindOne(ctx, bson.M{"_id": itemId}).Decode(&item)
	if err != nil {
		return InventoryItem{}, err
	}
	return item, nil
}

func List(ctx context.Context, condition bson.M) ([]InventoryItem, error) {
	_, db := getDB()
	cursor, err := db.Collection(INVENTORY_COLLECTION).Find(ctx, condition)
	if err != nil {
		return []InventoryItem{}, err
	}
	result := []InventoryItem{}
	err = cursor.All(ctx, &result)
	return result, err
}
