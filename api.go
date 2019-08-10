package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://Abhi_0512:12345ok@db-xhiqs.mongodb.net/test")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	router := mux.NewRouter()
	router.HandleFunc("/Person", CreatePerson).Methods("POST")
	//	router.HandleFunc("/peoples", GetPeopleEndpoint).Methods("GET")
	//	router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	http.ListenAndServe(":12345", router)
}

var client *mongo.Client

func CreatePerson(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("users").Collection("details")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}
