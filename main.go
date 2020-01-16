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
	http.HandleFunc("/updateContractData", mc.updateContractData)
	http.HandleFunc("/deleteContractData", mc.deleteContractData)
	http.HandleFunc("/selectContractData", mc.selectContractData)
	http.HandleFunc("/selectContractDataOne", mc.selectContractDataOne)

	http.HandleFunc("/insertDriverData", mc.insertDriverData)
	http.HandleFunc("/updateDriverData", mc.updateDriverData)
	http.HandleFunc("/deleteDriverData", mc.deleteDriverData)
	http.HandleFunc("/selectDriverData", mc.selectDriverData)
	http.HandleFunc("/selectDriverDataOne", mc.selectDriverDataOne)

	http.HandleFunc("/insertAutomobileData", mc.insertAutomobileData)
	http.HandleFunc("/updateAutomobileData", mc.updateAutomobileData)
	http.HandleFunc("/deleteAutomobileData", mc.deleteAutomobileData)
	http.HandleFunc("/selecAutomobileData", mc.selectAutomobileData)
	http.HandleFunc("/selectAutomobileDataOne", mc.selectAutomobileDataOne)

	http.HandleFunc("/insertTechnicalServiceData", mc.insertTechnicalServiceData)
	http.HandleFunc("/updateTechnicalServiceData", mc.updateTechnicalServiceData)
	http.HandleFunc("/deleteTechnicalServiceData", mc.deleteTechnicalServiceData)
	http.HandleFunc("/selectTechnicalServiceData", mc.selectTechnicalServiceData)
	http.HandleFunc("/selectTechnicalServiceDataOne", mc.selectTechnicalServiceDataOne)

	http.HandleFunc("/updateAccountData", mc.updateAccountData)
	http.HandleFunc("/insertAccountData", mc.insertAccountData)
	http.HandleFunc("/selectAccountData", mc.selectAccountData)
	http.HandleFunc("/deleteAccountData", mc.deleteAccountData)

	http.HandleFunc("/insertUserCompanyData", mc.insertUserCompanyData)
	http.HandleFunc("/saveUserCompanyData", mc.saveUserCompanyData)
	http.HandleFunc("/selectUserCompanyData", mc.selectUserCompanyDataOne)

	http.HandleFunc("/insertOwnerData", mc.insertOwnerData)
	http.HandleFunc("/updateOwnerData", mc.updateOwnerData)
	http.HandleFunc("/selectOwnerData", mc.selectOwnerData)
	http.HandleFunc("/selectAccountDataId", mc.selectAccountDataId)
	http.HandleFunc("/deleteOwnerData", mc.deleteOwnerData)
	panic(http.ListenAndServe(":8081", nil))
}
