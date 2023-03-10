package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/subhammahanty235/store-api-golang/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "storeApi"
const colName = "items"

var collection *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	connectionURI := os.Getenv("MONGODB_URI")

	clientOption := options.Client().ApplyURI(connectionURI)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(colName)
}

func insertNewItem(item model.Store) {
	inserted, err := collection.InsertOne(context.Background(), item)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted new Item", inserted.InsertedID)
}

func updateOneItem(id string, newData model.Store) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	updateData := make(map[string]interface{})
	if newData.ItemName != "" {
		updateData["itemname"] = newData.ItemName
	}
	if newData.Price > 0 {
		updateData["price"] = newData.Price
	}
	if newData.StockAvailable > 0 {
		updateData["stockavailable"] = newData.StockAvailable
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updateData}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("One item modified", )

}

func getAllItems() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var items []primitive.M
	for cursor.Next(context.Background()) {
		var item bson.M
		err := cursor.Decode(&item)

		if err != nil {
			log.Fatal(err)
		}

		items = append(items, item)

	}

	defer cursor.Close(context.Background())
	return items
}

// controllers

func InsertOneItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var item model.Store

	_ = json.NewDecoder(r.Body).Decode(&item)
	insertNewItem(item)

	json.NewEncoder(w).Encode(item)
}

func UpdateOneItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var newData model.Store
	_ = json.NewDecoder(r.Body).Decode(&newData)
	params := mux.Vars(r)

	updateOneItem(params["id"], newData)

	json.NewEncoder(w).Encode(params)

}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allItems := getAllItems()
	json.NewEncoder(w).Encode(allItems)
}
