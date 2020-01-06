package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

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

func (mc *MyClient) insertContractData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data Data
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	podcastsCollection := mc.db.Collection("contractsRent")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
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
		{"dateInsert", time.Now()},
		{"dateUpdate", nil},
	})
	if err != nil {
		log.Fatal(err)
	}
}
