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

type OverheadConsumableData struct {
	Id       string `json:"_id"`
	Type     string `json:"type"`
	Provider string `json:"provider"`
}

func (mc *MyClient) insertOverheadConsumableData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data OverheadConsumableData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	podcastsCollection := mc.db.Collection("overheadConsumable")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"type", data.Type},
		{"provider", data.Provider},
		{"dateInsert", time.Now()},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateOverheadConsumableData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data OverheadConsumableData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}
	podcastsCollection := mc.db.Collection("overheadConsumable")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resultUpdate, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"type":       data.Type,
				"provider":   data.Provider,
				"dateUpdate": time.Now(),
			},
		},
	)
	fmt.Println(resultUpdate.ModifiedCount) // output: 1
}

func (mc *MyClient) selectOverheadConsumableData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("overheadConsumable")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := podcastsCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData []OverheadConsumableData
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		idJson, err := json.Marshal(result["_id"])
		providerJson, err := json.Marshal(result["provider"])
		typeJson, err := json.Marshal(result["type"])

		idStr, _ := strconv.Unquote(string(idJson))
		providerStr, _ := strconv.Unquote(string(providerJson))
		typeStr, _ := strconv.Unquote(string(typeJson))

		parsedData = append(parsedData, OverheadConsumableData{
			Id:       string(idStr),
			Provider: string(providerStr),
			Type:     string(typeStr),
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))
}

func (mc *MyClient) selectOverheadConsumableDataOne(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var data OverheadConsumableData
	r.ParseForm()
	idGet := string(r.Form.Get("id"))
	id, err := primitive.ObjectIDFromHex(strings.Trim(idGet, "\""))
	if err != nil {
		fmt.Println(err)
	}

	podcastsCollection := mc.db.Collection("overheadConsumable")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.FindOne(
		ctx,
		bson.M{"_id": id}).Decode(&data)

	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)

	w.Write([]byte(bytes))
}

func (mc *MyClient) deleteOverheadConsumableData(w http.ResponseWriter, r *http.Request) {
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
	podcastsCollection := mc.db.Collection("overheadConsumable")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}
