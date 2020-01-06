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

type DataOwner struct {
	name               string `json:"name"`
	phone              string `json:"phone"`
	contactPerson      string `json:"contactPerson"`
	proceedings        string `json:"proceedings"`
	groundsForContract string `json:"groundsForContract"`
	additionContract   string `json:"additionContract"`
	percentageRevenue  string `json:"percentageRevenue"`
	profitInterest     string `json:"profitInterest"`
	perDay             string `json:"perDay"`
	perMounth          string `json:"perMounth"`
	conditionJobs      string `json:"conditionJobs"`
}

func (mc *MyClient) insertOwnerData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data DataOwner
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	podcastsCollection := mc.db.Collection("owners")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"name", data.name},
		{"phone", data.phone},
		{"contactPerson", data.contactPerson},
		{"proceedings", data.proceedings},
		{"groundsForContract", data.groundsForContract},
		{"additionContract", data.additionContract},
		{"percentageRevenue", data.percentageRevenue},
		{"profitInterest", data.profitInterest},
		{"perDay", data.perDay},
		{"perMounth", data.perMounth},
		{"conditionJobs", data.conditionJobs},
		{"dateInsert", time.Now()},
		{"dateUpdate", nil},
	})
	if err != nil {
		log.Fatal(err)
	}
}
