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

type AccountBillItemData struct {
	Id          string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

func (mc *MyClient) insertAccountBillItemData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data AccountBillItemData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	podcastsCollection := mc.db.Collection("accountBillItem")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"name", data.Name},
		{"description", data.Description},
		{"type", data.Type},
		{"dateInsert", time.Now()},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateAccountBillItemData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data AccountBillItemData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}
	podcastsCollection := mc.db.Collection("accountBillItem")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resultUpdate, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"name":        data.Name,
				"description": data.Description,
				"type":        data.Type,
				"dateUpdate":  time.Now(),
			},
		},
	)
	fmt.Println(resultUpdate.ModifiedCount) // output: 1
}

func (mc *MyClient) selectAccountBillItemData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("accountBillItem")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := podcastsCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData []AccountBillItemData
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		idJson, err := json.Marshal(result["_id"])
		nameJson, err := json.Marshal(result["name"])
		descriptionJson, err := json.Marshal(result["description"])
		typeJson, err := json.Marshal(result["type"])

		idStr, _ := strconv.Unquote(string(idJson))
		nameStr, _ := strconv.Unquote(string(nameJson))
		descriptionStr, _ := strconv.Unquote(string(descriptionJson))
		typeStr, _ := strconv.Unquote(string(typeJson))

		parsedData = append(parsedData, AccountBillItemData{
			Id:          string(idStr),
			Name:        string(nameStr),
			Description: string(descriptionStr),
			Type:        string(typeStr),
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))
}

func (mc *MyClient) selectAccountBillItemDataOne(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var data AccountBillItemData
	r.ParseForm()
	idGet := string(r.Form.Get("id"))
	id, err := primitive.ObjectIDFromHex(strings.Trim(idGet, "\""))
	if err != nil {
		fmt.Println(err)
	}

	podcastsCollection := mc.db.Collection("accountBillItem")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.FindOne(
		ctx,
		bson.M{"_id": id}).Decode(&data)

	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)

	w.Write([]byte(bytes))
}

func (mc *MyClient) deleteAccountBillItemData(w http.ResponseWriter, r *http.Request) {
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
	podcastsCollection := mc.db.Collection("accountBillItem")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}
