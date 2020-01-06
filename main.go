package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MyClient struct {
	mc *mongo.Client
	db *mongo.Database
}

func setupResponse(w http.ResponseWriter, req *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func NewMyClient(url, db string) (mc *MyClient, err error) {
	mc = &MyClient{}
	if mc.mc, err = mongo.NewClient(options.Client().ApplyURI(url)); err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = mc.mc.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	mc.db = mc.mc.Database(db)
	return
}

func main() {
	mc, err := NewMyClient("mongodb://localhost:27017", "crmTaxi")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/insertContractData", mc.insertContractData)
	http.HandleFunc("/insertDriverData", mc.insertDriverData)
	http.HandleFunc("/insertAutomobileData", mc.insertAutomobileData)
	http.HandleFunc("/insertTechnicalServiceData", mc.insertTechnicalServiceData)
	http.HandleFunc("/insertOwnerData", mc.insertOwnerData)
	panic(http.ListenAndServe(":8081", nil))
}
