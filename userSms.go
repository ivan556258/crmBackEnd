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

type UserSms struct {
	Id        string `json:"_id"`
	Recipient string `json:"recipient"`
	Phone     string `json:"phone"`
	Message   string `json:"message"`
	Token     string `json:"token"`
}

func (mc *MyClient) insertUserSmsData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data UserSms
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	podcastsCollection := mc.db.Collection("userSms")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"recipient", data.Recipient},
		{"phone", data.Phone},
		{"token", data.Token},
		{"message", data.Message},
		{"dateUpdate", time.Now()},
		{"dateInsert", time.Now()},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateUserSmsData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data UserSms
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}
	podcastsCollection := mc.db.Collection("userSms")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resultUpdate, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"recipient":  data.Recipient,
				"phone":      data.Phone,
				"message":    data.Message,
				"dateUpdate": time.Now(),
			},
		},
	)
	fmt.Println(resultUpdate.ModifiedCount) // output: 1
}

func (mc *MyClient) deleteUserSmsData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data UserSms
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
	podcastsCollection := mc.db.Collection("userSms")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}

func (mc *MyClient) selectUserSmsData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("userSms")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.ParseForm()
	token := string(r.Form.Get("token"))
	cur, err := podcastsCollection.Find(ctx, bson.D{{"token", token}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData []UserSms
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		idJson, err := json.Marshal(result["_id"])
		recipientJson, err := json.Marshal(result["recipient"])
		phoneJson, err := json.Marshal(result["phone"])
		messageJson, err := json.Marshal(result["message"])

		idStr, _ := strconv.Unquote(string(idJson))
		recipientStr, _ := strconv.Unquote(string(recipientJson))
		phoneStr, _ := strconv.Unquote(string(phoneJson))
		messageStr, _ := strconv.Unquote(string(messageJson))

		parsedData = append(parsedData, UserSms{
			Id:        string(idStr),
			Recipient: string(recipientStr),
			Phone:     string(phoneStr),
			Message:   string(messageStr),
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))
}
