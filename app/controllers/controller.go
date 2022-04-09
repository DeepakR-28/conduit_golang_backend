package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	database "github.com/deepakr-28/conduit_golang_backend/app/database"
	model "github.com/deepakr-28/conduit_golang_backend/app/models"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection

// const dbName = "conduit_golang_backend"
const collectionName = "users"

var databaseName string

func insertUser(user model.User) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseName = os.Getenv("MONGODB_CONNECTION_STRING")
	res, err := database.Client.ListDatabaseNames(context.Background(), bson.M{})
	// defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	Collection = database.Client.Database(databaseName).Collection(collectionName)

	inserted, err := Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 user in db with id: ", inserted.InsertedID)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user model.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	insertUser(user)

	json.NewEncoder(w).Encode(user)
}
