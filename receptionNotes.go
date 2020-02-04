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

type ReceptionNoteData struct {
	Id       string `json:"_id"`
	Type     string `json:"type"`
	Provider string `json:"provider"`
	Token    string `json:"token"`
}

func (mc *MyClient) insertReceptionNoteData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data ReceptionNoteData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	podcastsCollection := mc.db.Collection("receptionNotes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"type", data.Type},
		{"token", data.Token},
		{"provider", data.Provider},
		{"dateInsert", time.Now()},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateReceptionNoteData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data ReceptionNoteData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}
	podcastsCollection := mc.db.Collection("receptionNotes")
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

func (mc *MyClient) selectReceptionNoteData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("receptionNotes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.ParseForm()
	token := string(r.Form.Get("token"))
	cur, err := podcastsCollection.Find(ctx, bson.D{{"token", token}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData []ReceptionNoteData
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

		parsedData = append(parsedData, ReceptionNoteData{
			Id:       string(idStr),
			Provider: string(providerStr),
			Type:     string(typeStr),
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))
}

func (mc *MyClient) selectReceptionNoteDataOne(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var data ReceptionNoteData
	r.ParseForm()
	idGet := string(r.Form.Get("id"))
	id, err := primitive.ObjectIDFromHex(strings.Trim(idGet, "\""))
	if err != nil {
		fmt.Println(err)
	}

	podcastsCollection := mc.db.Collection("receptionNotes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.FindOne(
		ctx,
		bson.M{"_id": id}).Decode(&data)

	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)

	w.Write([]byte(bytes))
}

func (mc *MyClient) deleteReceptionNoteData(w http.ResponseWriter, r *http.Request) {
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
	podcastsCollection := mc.db.Collection("receptionNotes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}
