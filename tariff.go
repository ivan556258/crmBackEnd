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

type TariffData struct {
	Id               string `json:"_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	CategoryAuto     string `json:"categoryAuto"`
	Tariff           string `json:"tariff"`
	Network          string `json:"network"`
	CoastPerDay      string `json:"coastPerDay"`
	ContractContinue string `json:"contractContinue"`
	StartPayment     string `json:"startPayment"`
	SummAmount       string `json:"summAmount"`
	StatusRes        string `json:"statusRes"`
	Token            string `json:"token"`
	SummPerMounth    string `json:"summPerMounth"`
	SummPerDay       string `json:"summPerDay"`
	SummMo           string `json:"summMo"`
	SummTu           string `json:"summTu"`
	SummWe           string `json:"summWe"`
	SummTh           string `json:"summTh"`
	SummFr           string `json:"summFr"`
	SummSa           string `json:"summSa"`
	SummSu           string `json:"summSu"`
}

func (mc *MyClient) insertTariffData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data TariffData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	podcastsCollection := mc.db.Collection("tariff")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"name", data.Name},
		{"description", data.Description},
		{"categoryAuto", data.CategoryAuto},
		{"tariff", data.Tariff},
		{"network", data.Network},
		{"coastPerDay", data.CoastPerDay},
		{"contractContinue", data.ContractContinue},
		{"startPayment", data.StartPayment},
		{"summAmount", data.SummAmount},
		{"summPerMounth", data.SummPerMounth},
		{"summPerDay", data.SummPerDay},
		{"summMo", data.SummMo},
		{"summTu", data.SummTu},
		{"summWe", data.SummWe},
		{"summTh", data.SummTh},
		{"summFr", data.SummFr},
		{"summSa", data.SummSa},
		{"summSu", data.SummSu},
		{"token", data.Token},
		{"statusRes", data.StatusRes},
		{"dateInsert", time.Now()},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateTariffData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data TariffData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}
	podcastsCollection := mc.db.Collection("tariff")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resultUpdate, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"name":             data.Name,
				"description":      data.Description,
				"categoryAuto":     data.CategoryAuto,
				"tariff":           data.Tariff,
				"network":          data.Network,
				"coastPerDay":      data.CoastPerDay,
				"contractContinue": data.ContractContinue,
				"startPayment":     data.StartPayment,
				"summAmount":       data.SummAmount,
				"summPerMounth":    data.SummPerMounth,
				"summPerDay":       data.SummPerDay,
				"summMo":           data.SummMo,
				"summTu":           data.SummTu,
				"summWe":           data.SummWe,
				"summTh":           data.SummTh,
				"summFr":           data.SummFr,
				"summSa":           data.SummSa,
				"summSu":           data.SummSu,
				"statusRes":        data.StatusRes,
				"dateUpdate":       time.Now(),
			},
		},
	)
	fmt.Println(resultUpdate.ModifiedCount) // output: 1
}

func (mc *MyClient) selectTariffData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("tariff")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.ParseForm()
	token := string(r.Form.Get("token"))
	cur, err := podcastsCollection.Find(ctx, bson.D{{"token", token}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData []TariffData
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		idJson, err := json.Marshal(result["_id"])
		nameJson, err := json.Marshal(result["name"])
		descriptionJson, err := json.Marshal(result["description"])
		categoryAutoJson, err := json.Marshal(result["categoryAuto"])
		tariffJson, err := json.Marshal(result["tariff"])
		networkJson, err := json.Marshal(result["network"])
		coastPerDayJson, err := json.Marshal(result["coastPerDay"])
		contractContinueJson, err := json.Marshal(result["contractContinue"])
		startPaymentJson, err := json.Marshal(result["startPayment"])
		summAmountJson, err := json.Marshal(result["summAmount"])
		statusResJson, err := json.Marshal(result["statusRes"])
		summPerMounthJson, err := json.Marshal(result["summPerMounth"])
		summPerDayJson, err := json.Marshal(result["summPerDay"])
		summMoJson, err := json.Marshal(result["summMo"])
		summTuJson, err := json.Marshal(result["summTu"])
		summWeJson, err := json.Marshal(result["summWe"])
		summThJson, err := json.Marshal(result["summTh"])
		summFrJson, err := json.Marshal(result["summFr"])
		summSaJson, err := json.Marshal(result["summSa"])
		summSuJson, err := json.Marshal(result["summSu"])

		idStr, _ := strconv.Unquote(string(idJson))
		nameStr, _ := strconv.Unquote(string(nameJson))
		descriptionStr, _ := strconv.Unquote(string(descriptionJson))
		categoryAutoStr, _ := strconv.Unquote(string(categoryAutoJson))
		tariffStr, _ := strconv.Unquote(string(tariffJson))
		networkStr, _ := strconv.Unquote(string(networkJson))
		coastPerDayStr, _ := strconv.Unquote(string(coastPerDayJson))
		contractContinueStr, _ := strconv.Unquote(string(contractContinueJson))
		startPaymentStr, _ := strconv.Unquote(string(startPaymentJson))
		summAmountStr, _ := strconv.Unquote(string(summAmountJson))
		statusResStr, _ := strconv.Unquote(string(statusResJson))
		summPerMounthStr, _ := strconv.Unquote(string(summPerMounthJson))
		summPerDayStr, _ := strconv.Unquote(string(summPerDayJson))
		summMoStr, _ := strconv.Unquote(string(summMoJson))
		summTuStr, _ := strconv.Unquote(string(summTuJson))
		summWeStr, _ := strconv.Unquote(string(summWeJson))
		summThStr, _ := strconv.Unquote(string(summThJson))
		summFrStr, _ := strconv.Unquote(string(summFrJson))
		summSaStr, _ := strconv.Unquote(string(summSaJson))
		summSuStr, _ := strconv.Unquote(string(summSuJson))

		parsedData = append(parsedData, TariffData{
			Id:               string(idStr),
			Name:             string(nameStr),
			Description:      string(descriptionStr),
			CategoryAuto:     string(categoryAutoStr),
			Tariff:           string(tariffStr),
			Network:          string(networkStr),
			CoastPerDay:      string(coastPerDayStr),
			ContractContinue: string(contractContinueStr),
			StartPayment:     string(startPaymentStr),
			SummAmount:       string(summAmountStr),
			StatusRes:        string(statusResStr),
			SummPerMounth:    string(summPerMounthStr),
			SummPerDay:       string(summPerDayStr),
			SummMo:           string(summMoStr),
			SummTu:           string(summTuStr),
			SummWe:           string(summWeStr),
			SummTh:           string(summThStr),
			SummFr:           string(summFrStr),
			SummSa:           string(summSaStr),
			SummSu:           string(summSuStr),
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))
}

func (mc *MyClient) selectTariffDataOne(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var data TariffData
	r.ParseForm()
	idGet := string(r.Form.Get("id"))
	id, err := primitive.ObjectIDFromHex(strings.Trim(idGet, "\""))
	if err != nil {
		fmt.Println(err)
	}

	podcastsCollection := mc.db.Collection("tariff")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.FindOne(
		ctx,
		bson.M{"_id": id}).Decode(&data)

	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)

	w.Write([]byte(bytes))
}

func (mc *MyClient) deleteTariffData(w http.ResponseWriter, r *http.Request) {
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
	podcastsCollection := mc.db.Collection("tariff")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}
