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

type NomenclatureData struct {
	Id         string `json:"_id"`
	Name       string `json:"name"`
	Brand      string `json:"brand"`
	Article    string `json:"article"`
	IsPart     bool   `json:"isPart"`
	Summ       string `json:"summ"`
	HowMuch    string `json:"howMuch"`
	Token      string `json:"token"`
	DateInsert string `json:"dateInsert"`
}

func (mc *MyClient) insertNomenclatureData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data NomenclatureData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	podcastsCollection := mc.db.Collection("nomenclature")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"name", data.Name},
		{"brand", data.Brand},
		{"article", data.Article},
		{"isPart", data.IsPart},
		{"summ", data.Summ},
		{"howMuch", data.HowMuch},
		{"token", data.Token},
		{"dateInsert", time.Now()},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateNomenclatureData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data NomenclatureData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}
	podcastsCollection := mc.db.Collection("nomenclature")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resultUpdate, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"name":       data.Name,
				"brand":      data.Brand,
				"article":    data.Article,
				"isPart":     data.IsPart,
				"summ":       data.Summ,
				"howMuch":    data.HowMuch,
				"dateUpdate": time.Now(),
			},
		},
	)
	fmt.Println(resultUpdate.ModifiedCount) // output: 1
}

func (mc *MyClient) selectNomenclatureData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("nomenclature")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.ParseForm()
	token := string(r.Form.Get("token"))
	cur, err := podcastsCollection.Find(ctx, bson.D{{"token", token}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData []NomenclatureData
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		idJson, err := json.Marshal(result["_id"])
		nameJson, err := json.Marshal(result["name"])
		brandJson, err := json.Marshal(result["brand"])
		articleJson, err := json.Marshal(result["article"])
		isPartJson, err := json.Marshal(result["isPart"])
		summJson, err := json.Marshal(result["summ"])
		howMuchJson, err := json.Marshal(result["howMuch"])
		dateUpdateJson, err := json.Marshal(result["dateUpdate"])

		idStr, _ := strconv.Unquote(string(idJson))
		nameStr, _ := strconv.Unquote(string(nameJson))
		brandStr, _ := strconv.Unquote(string(brandJson))
		articleStr, _ := strconv.Unquote(string(articleJson))
		isPartStr, _ := strconv.ParseBool(string(isPartJson))
		summStr, _ := strconv.Unquote(string(summJson))
		howMuchStr, _ := strconv.Unquote(string(howMuchJson))
		dateUpdateStr, _ := strconv.Unquote(string(dateUpdateJson))

		parsedData = append(parsedData, NomenclatureData{
			Id:         string(idStr),
			Name:       string(nameStr),
			Brand:      string(brandStr),
			Article:    string(articleStr),
			IsPart:     bool(isPartStr),
			Summ:       string(summStr),
			HowMuch:    string(howMuchStr),
			DateInsert: string(dateUpdateStr),
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))
}

func (mc *MyClient) selectNomenclatureIsPlaceData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("nomenclature")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.ParseForm()
	token := string(r.Form.Get("token"))
	cur, err := podcastsCollection.Find(ctx, bson.D{{"token", token}, {"isPart", true}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData []NomenclatureData
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		idJson, err := json.Marshal(result["_id"])
		nameJson, err := json.Marshal(result["name"])
		brandJson, err := json.Marshal(result["brand"])
		articleJson, err := json.Marshal(result["article"])
		isPartJson, err := json.Marshal(result["isPart"])
		summJson, err := json.Marshal(result["summ"])
		howMuchJson, err := json.Marshal(result["howMuch"])
		dateUpdateJson, err := json.Marshal(result["dateUpdate"])

		idStr, _ := strconv.Unquote(string(idJson))
		nameStr, _ := strconv.Unquote(string(nameJson))
		brandStr, _ := strconv.Unquote(string(brandJson))
		articleStr, _ := strconv.Unquote(string(articleJson))
		isPartStr, _ := strconv.ParseBool(string(isPartJson))
		summStr, _ := strconv.Unquote(string(summJson))
		howMuchStr, _ := strconv.Unquote(string(howMuchJson))
		dateUpdateStr, _ := strconv.Unquote(string(dateUpdateJson))

		parsedData = append(parsedData, NomenclatureData{
			Id:         string(idStr),
			Name:       string(nameStr),
			Brand:      string(brandStr),
			Article:    string(articleStr),
			IsPart:     bool(isPartStr),
			Summ:       string(summStr),
			HowMuch:    string(howMuchStr),
			DateInsert: string(dateUpdateStr),
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))
}

func (mc *MyClient) selectNomenclatureDataOne(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var data NomenclatureData
	r.ParseForm()
	idGet := string(r.Form.Get("id"))
	id, err := primitive.ObjectIDFromHex(strings.Trim(idGet, "\""))
	if err != nil {
		fmt.Println(err)
	}

	podcastsCollection := mc.db.Collection("nomenclature")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.FindOne(
		ctx,
		bson.M{"_id": id}).Decode(&data)

	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)

	w.Write([]byte(bytes))
}

func (mc *MyClient) deleteNomenclatureData(w http.ResponseWriter, r *http.Request) {
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
	podcastsCollection := mc.db.Collection("nomenclature")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}
