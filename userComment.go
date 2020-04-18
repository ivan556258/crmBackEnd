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

type UserComment struct {
	Id      string `json:"_id"`
	Driver  bool   `json:"driver"`
	Comment string `json:"comment"`
	Notify  string `json:"notify"`
	Token   string `json:"token"`
}

func (mc *MyClient) insertUserCommentData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data UserComment
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	podcastsCollection := mc.db.Collection("userComment")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"driver", data.Driver},
		{"comment", data.Comment},
		{"notify", data.Notify},
		{"token", data.Token},
		{"dateUpdate", time.Now()},
		{"dateInsert", time.Now()},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateUserCommentData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data UserComment
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}
	podcastsCollection := mc.db.Collection("userComment")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"driver":     data.Driver,
				"comment":    data.Comment,
				"notify":     data.Notify,
				"dateUpdate": time.Now(),
			},
		},
	)
	//fmt.Println(resultUpdate.ModifiedCount) // output: 1
}
func (mc *MyClient) deleteUserCommentData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data UserComment
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}
	podcastsCollection := mc.db.Collection("userComment")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}

func (mc *MyClient) selectUserCommentData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("userComment")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.ParseForm()
	token := string(r.Form.Get("token"))
	cur, err := podcastsCollection.Find(ctx, bson.D{{"token", token}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData []UserComment
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		idJson, err := json.Marshal(result["_id"])
		driverJson, err := json.Marshal(result["driver"])
		commentJson, err := json.Marshal(result["comment"])
		notifyJson, err := json.Marshal(result["notify"])

		idStr, _ := strconv.Unquote(string(idJson))
		driverStr, _ := strconv.ParseBool(string(driverJson))
		commentStr, _ := strconv.Unquote(string(commentJson))
		notifyStr, _ := strconv.Unquote(string(notifyJson))

		parsedData = append(parsedData, UserComment{
			Id:      string(idStr),
			Driver:  bool(driverStr),
			Comment: string(commentStr),
			Notify:  string(notifyStr),
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))
}
