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

type AccountBillData struct {
	Id             string `json:"_id"`
	Cash           string `json:"cash"`
	Item           string `json:"item"`
	Summ           string `json:"summ"`
	Description    string `json:"description"`
	Protein        string `json:"protein"`
	AddTransaction bool   `json:"addTransaction"`
	StatusCash     int    `json:"statusCash"` // если списываем
}

func (mc *MyClient) insertAccountBillData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data AccountBillData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	podcastsCollection := mc.db.Collection("accountBill")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"cash", data.Cash},
		{"item", data.Item},
		{"summ", data.Summ},
		{"description", data.Description},
		{"protein", data.Protein},
		{"addTransaction", data.AddTransaction},
		{"statusCash", data.StatusCash}, // если списываем
		{"dateInsert", time.Now()},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateAccountBillData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data AccountBillData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}
	podcastsCollection := mc.db.Collection("accountBill")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resultUpdate, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"cash":           data.Cash,
				"item":           data.Item,
				"summ":           data.Summ,
				"description":    data.Description,
				"protein":        data.Protein,
				"addTransaction": data.AddTransaction,
				"statusCash":     data.StatusCash, // если списываем
				"dateUpdate":     time.Now(),
			},
		},
	)
	fmt.Println(resultUpdate.ModifiedCount) // output: 1
}

func (mc *MyClient) selectAccountBillData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("accountBill")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := podcastsCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData []AccountBillData
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		idJson, err := json.Marshal(result["_id"])
		cashJson, err := json.Marshal(result["cash"])
		itemJson, err := json.Marshal(result["item"])
		summJson, err := json.Marshal(result["summ"])
		descriptionJson, err := json.Marshal(result["description"])
		proteinJson, err := json.Marshal(result["protein"])
		addTransactionJson, err := json.Marshal(result["addTransaction"])
		statusCashJson, err := json.Marshal(result["statusCash"]) // если списываем

		idStr, _ := strconv.Unquote(string(idJson))
		cashStr, _ := strconv.Unquote(string(cashJson))
		itemStr, _ := strconv.Unquote(string(itemJson))
		summStr, _ := strconv.Unquote(string(summJson))
		descriptionStr, _ := strconv.Unquote(string(descriptionJson))
		proteinStr, _ := strconv.Unquote(string(proteinJson))
		addTransactionStr, _ := strconv.ParseBool(string(addTransactionJson))
		statusCashStr, _ := strconv.Atoi(string(statusCashJson)) // если списываем

		parsedData = append(parsedData, AccountBillData{
			Id:             string(idStr),
			Cash:           string(cashStr),
			Item:           string(itemStr),
			Summ:           string(summStr),
			Description:    string(descriptionStr),
			Protein:        string(proteinStr),
			AddTransaction: bool(addTransactionStr),
			StatusCash:     int(statusCashStr), // если списываем
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))
}

func (mc *MyClient) selectAccountBillDataOne(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var data AccountBillData
	r.ParseForm()
	idGet := string(r.Form.Get("id"))
	id, err := primitive.ObjectIDFromHex(strings.Trim(idGet, "\""))
	if err != nil {
		fmt.Println(err)
	}

	podcastsCollection := mc.db.Collection("accountBill")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.FindOne(
		ctx,
		bson.M{"_id": id}).Decode(&data)

	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)

	w.Write([]byte(bytes))
}

func (mc *MyClient) deleteAccountBillData(w http.ResponseWriter, r *http.Request) {
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
	podcastsCollection := mc.db.Collection("accountBill")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}
