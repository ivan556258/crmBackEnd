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

type CounteragentData struct {
	Id            string `json:"_id"`
	Name          string `json:"name"`
	Website       string `json:"website"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Type          string `json:"type"`
	NameLow       string `json:"nameLow"`
	Address       string `json:"address"`
	ContactPerson string `json:"contactPerson"`
	Note          string `json:"note"`
	Token         string `json:"token"`
}

func (mc *MyClient) insertCounteragentData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data CounteragentData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	podcastsCollection := mc.db.Collection("counteragents")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"name", data.Name},
		{"website", data.Website},
		{"phone", data.Phone},
		{"email", data.Email},
		{"type", data.Type},
		{"token", data.Token},
		{"nameLow", data.NameLow},
		{"address", data.Address},
		{"contactPerson", data.ContactPerson},
		{"note", data.Note},
		{"dateInsert", time.Now()},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateCounteragentData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data CounteragentData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}
	podcastsCollection := mc.db.Collection("counteragents")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resultUpdate, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"name":          data.Name,
				"website":       data.Website,
				"phone":         data.Phone,
				"email":         data.Email,
				"type":          data.Type,
				"nameLow":       data.NameLow,
				"address":       data.Address,
				"contactPerson": data.ContactPerson,
				"note":          data.Note,
				"dateUpdate":    time.Now(),
			},
		},
	)
	fmt.Println(resultUpdate.ModifiedCount) // output: 1
}

func (mc *MyClient) selectCounteragentData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("counteragents")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.ParseForm()
	token := string(r.Form.Get("token"))
	cur, err := podcastsCollection.Find(ctx, bson.D{{"token", token}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData []CounteragentData
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		idJson, err := json.Marshal(result["_id"])
		nameJson, err := json.Marshal(result["name"])
		websiteJson, err := json.Marshal(result["website"])
		phoneJson, err := json.Marshal(result["phone"])
		emailJson, err := json.Marshal(result["email"])
		typeJson, err := json.Marshal(result["type"])
		nameLowJson, err := json.Marshal(result["nameLow"])
		addressJson, err := json.Marshal(result["address"])
		contactPersonJson, err := json.Marshal(result["contactPerson"])
		noteJson, err := json.Marshal(result["note"])

		idStr, _ := strconv.Unquote(string(idJson))
		nameStr, _ := strconv.Unquote(string(nameJson))
		websiteStr, _ := strconv.Unquote(string(websiteJson))
		phoneStr, _ := strconv.Unquote(string(phoneJson))
		emailStr, _ := strconv.Unquote(string(emailJson))
		typeStr, _ := strconv.Unquote(string(typeJson))
		nameLowStr, _ := strconv.Unquote(string(nameLowJson))
		addressStr, _ := strconv.Unquote(string(addressJson))
		contactPersonStr, _ := strconv.Unquote(string(contactPersonJson))
		noteStr, _ := strconv.Unquote(string(noteJson))

		parsedData = append(parsedData, CounteragentData{
			Id:            string(idStr),
			Name:          string(nameStr),
			Website:       string(websiteStr),
			Phone:         string(phoneStr),
			Email:         string(emailStr),
			Type:          string(typeStr),
			NameLow:       string(nameLowStr),
			Address:       string(addressStr),
			ContactPerson: string(contactPersonStr),
			Note:          string(noteStr),
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))
}

func (mc *MyClient) selectCounteragentDataOne(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var data CounteragentData
	r.ParseForm()
	idGet := string(r.Form.Get("id"))
	id, err := primitive.ObjectIDFromHex(strings.Trim(idGet, "\""))
	if err != nil {
		fmt.Println(err)
	}

	podcastsCollection := mc.db.Collection("counteragents")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.FindOne(
		ctx,
		bson.M{"_id": id}).Decode(&data)

	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)

	w.Write([]byte(bytes))
}

func (mc *MyClient) deleteCounteragentData(w http.ResponseWriter, r *http.Request) {
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
	podcastsCollection := mc.db.Collection("counteragents")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}
