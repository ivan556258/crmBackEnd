package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// etc

type Data struct {
	Number    string `json:"number"`
	Driver    string `json:"driver"`
	Auto      string `json:"auto"`
	Tariff    string `json:"tariff"`
	Begindate string `json:"begindate"`
	Enddate   string `json:"enddate"`
	Continues bool   `json:"continues"`
	MoreInfo  string `json:"moreInfo"`
	Status    string `json:"status"`
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func selectContractData(w http.ResponseWriter, r *http.Request) {

}

func insertContractData(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return
	}

	var (
		mongoURL = "mongodb://localhost:27017"
	)
	// Initialize a new mongo client with options
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	quickstartDatabase := client.Database("new_databse")
	podcastsCollection := quickstartDatabase.Collection("test")
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"number", data.Number},
		{"driver", data.Driver},
		{"auto", data.Auto},
		{"tariff", data.Tariff},
		{"begindate", data.Begindate},
		{"enddate", data.Enddate},
		{"continues", data.Continues},
		{"moreInfo", data.MoreInfo},
		{"status", data.Status},
	})
	if err != nil {
		log.Fatal(err)
	}

}

func ConnectMongo() {
	/* var (
		mongoURL = "mongodb://localhost:27017"
	)
	// Initialize a new mongo client with options
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx) */

}

func main() {
	http.HandleFunc("/insertContractData", insertContractData)
	http.HandleFunc("/selectContractData", selectContractData)
	http.ListenAndServe(":8081", nil)
}
