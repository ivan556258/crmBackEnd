package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionData struct {
	Id             string `json:"_id"`
	Item           string `json:"item"`
	Score          string `json:"score"`
	Summ           string `json:"summ"`
	Description    string `json:"description"`
	AddTransaction bool   `json:"addTransaction"`
	Name           string `json:"name"`
	Tariff         string `json:"tariff"`
	Token          string `json:"token"`
	DateInsert     string `json:"dateInsert"`
	Autor          string `json:"autor"`
}

func (mc *MyClient) insertTransactionData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data TransactionData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	var name string = mc.dataNameFind("drivers", "_id", data.Score)
	var tariff string = mc.dataAccountBillItemFind("accountBillItem", "_id", data.Item)
	podcastsCollection := mc.db.Collection("transacrions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"item", data.Item},
		{"score", data.Score},
		{"summ", data.Summ},
		{"token", data.Token},
		{"description", data.Description},
		{"addTransaction", data.AddTransaction},
		{"name", name},
		{"tariff", tariff},
		{"dateInsert", time.Now()},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateTransactionData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data TransactionData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	var name string = mc.dataNameFind("drivers", "_id", data.Score)
	var tariff string = mc.dataAccountBillItemFind("accountBillItem", "_id", data.Item)
	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}
	podcastsCollection := mc.db.Collection("transacrions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resultUpdate, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"item":           data.Item,
				"score":          data.Score,
				"summ":           data.Summ,
				"description":    data.Description,
				"addTransaction": data.AddTransaction,
				"name":           name,
				"tariff":         tariff,
				"dateUpdate":     time.Now(),
			},
		},
	)
	fmt.Println(resultUpdate.ModifiedCount) // output: 1
}

func (mc *MyClient) selectTransactionData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("transacrions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.ParseForm()
	token := string(r.Form.Get("token"))
	cur, err := podcastsCollection.Find(ctx, bson.D{{"token", token}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData []TransactionData
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		idJson, err := json.Marshal(result["_id"])
		itemJson, err := json.Marshal(result["item"])
		scoreJson, err := json.Marshal(result["score"])
		summJson, err := json.Marshal(result["summ"])
		nameJson, err := json.Marshal(result["name"])
		tariffJson, err := json.Marshal(result["tariff"])
		descriptionJson, err := json.Marshal(result["description"])
		addTransactionJson, err := json.Marshal(result["addTransaction"])
		dateInsertJson, err := json.Marshal(result["dateInsert"])

		idStr, _ := strconv.Unquote(string(idJson))
		itemStr, _ := strconv.Unquote(string(itemJson))
		scoreStr, _ := strconv.Unquote(string(scoreJson))
		summStr, _ := strconv.Unquote(string(summJson))
		nameStr, _ := strconv.Unquote(string(nameJson))
		tariffStr, _ := strconv.Unquote(string(tariffJson))
		descriptionStr, _ := strconv.Unquote(string(descriptionJson))
		addTransactionStr, _ := strconv.ParseBool(string(addTransactionJson))
		dateInsertStr, _ := strconv.Unquote(string(dateInsertJson))

		parsedData = append(parsedData, TransactionData{
			Id:             string(idStr),
			Item:           string(itemStr),
			Score:          string(scoreStr),
			Summ:           string(summStr),
			Name:           string(nameStr),
			Tariff:         string(tariffStr),
			Description:    string(descriptionStr),
			AddTransaction: bool(addTransactionStr),
			DateInsert:     string(dateInsertStr),
			Autor:          string("ООО «АВТО ДРАЙВ»"),
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))
}

func (mc *MyClient) selectTransactionDataOne(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var data TransactionData
	r.ParseForm()
	idGet := string(r.Form.Get("id"))
	id, err := primitive.ObjectIDFromHex(strings.Trim(idGet, "\""))
	if err != nil {
		fmt.Println(err)
	}

	podcastsCollection := mc.db.Collection("transacrions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.FindOne(
		ctx,
		bson.M{"_id": id}).Decode(&data)

	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)

	w.Write([]byte(bytes))
}

func (mc *MyClient) deleteTransactionData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data DataAccount
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data.Id)
	podcastsCollection := mc.db.Collection("transacrions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}
