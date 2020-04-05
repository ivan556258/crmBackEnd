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

type DataOwner struct {
	Id                 string `json:"_id"`
	Name               string `json:"name"`
	Phone              string `json:"phone"`
	ContactPerson      string `json:"contactPerson"`
	Proceedings        string `json:"proceedings"`
	GroundsForContract string `json:"groundsForContract"`
	AdditionContract   string `json:"additionContract"`
	PercentageRevenue  string `json:"percentageRevenue"`
	ProfitInterest     string `json:"profitInterest"`
	PerDay             string `json:"perDay"`
	PerMounth          string `json:"perMounth"`
	ConditionJobs      string `json:"conditionJobs"`
	Token              string `json:"token"`
}

func (mc *MyClient) insertOwnerData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var data DataOwner
	var err error
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	podcastsCollection := mc.db.Collection("owners")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"name", data.Name},
		{"phone", data.Phone},
		{"contactPerson", data.ContactPerson},
		{"proceedings", data.Proceedings},
		{"groundsForContract", data.GroundsForContract},
		{"additionContract", data.AdditionContract},
		{"percentageRevenue", data.PercentageRevenue},
		{"profitInterest", data.ProfitInterest},
		{"perDay", data.PerDay},
		{"token", data.Token},
		{"free", 1},
		{"perMounth", data.PerMounth},
		{"conditionJobs", data.ConditionJobs},
		{"dateInsert", time.Now()},
		{"dateUpdate", nil},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateOwnerData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var data DataOwner
	var err error
	err = json.NewDecoder(r.Body).Decode(&data)
	fmt.Println(data.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}
	podcastsCollection := mc.db.Collection("owners")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resultUpdate, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"name":               data.Name,
				"phone":              data.Phone,
				"contactPerson":      data.ContactPerson,
				"proceedings":        data.Proceedings,
				"groundsForContract": data.GroundsForContract,
				"additionContract":   data.AdditionContract,
				"percentageRevenue":  data.PercentageRevenue,
				"profitInterest":     data.ProfitInterest,
				"perDay":             data.PerDay,
				"perMounth":          data.PerMounth,
				"conditionJobs":      data.ConditionJobs,
				"dateUpdate":         time.Now(),
			},
		},
	)
	fmt.Println(resultUpdate.ModifiedCount) // output: 1
}

func (mc *MyClient) selectOwnerData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var parsedData []DataOwner
	var err error
	podcastsCollection := mc.db.Collection("owners")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.ParseForm()
	token := string(r.Form.Get("token"))
	cur, err := podcastsCollection.Find(ctx, bson.D{{"token", token}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		id, err := json.Marshal(result["_id"])
		nameJson, err := json.Marshal(result["name"])
		phoneJson, err := json.Marshal(result["phone"])
		contactPersonJson, err := json.Marshal(result["contactPerson"])
		proceedingsJson, err := json.Marshal(result["proceedings"])
		groundsForContractJson, err := json.Marshal(result["groundsForContract"])
		additionContractJson, err := json.Marshal(result["additionContract"])
		percentageRevenueJson, err := json.Marshal(result["percentageRevenue"])
		profitInterestJson, err := json.Marshal(result["profitInterest"])
		perDayJson, err := json.Marshal(result["perDay"])
		perMounthJson, err := json.Marshal(result["perMounth"])
		conditionJobsJson, err := json.Marshal(result["conditionJobs"])

		sid, _ := strconv.Unquote(string(id))
		nameStr, _ := strconv.Unquote(string(nameJson))
		phoneStr, _ := strconv.Unquote(string(phoneJson))
		contactPersonStr, _ := strconv.Unquote(string(contactPersonJson))
		proceedingsStr, _ := strconv.Unquote(string(proceedingsJson))
		groundsForContractStr, _ := strconv.Unquote(string(groundsForContractJson))
		additionContractStr, _ := strconv.Unquote(string(additionContractJson))
		percentageRevenueStr, _ := strconv.Unquote(string(percentageRevenueJson))
		profitInterestStr, _ := strconv.Unquote(string(profitInterestJson))
		perDayStr, _ := strconv.Unquote(string(perDayJson))
		perMounthStr, _ := strconv.Unquote(string(perMounthJson))
		conditionJobsStr, _ := strconv.Unquote(string(conditionJobsJson))

		parsedData = append(parsedData, DataOwner{
			Id:                 string(sid),
			Name:               string(nameStr),
			Phone:              string(phoneStr),
			ContactPerson:      string(contactPersonStr),
			Proceedings:        string(proceedingsStr),
			GroundsForContract: string(groundsForContractStr),
			AdditionContract:   string(additionContractStr),
			PercentageRevenue:  string(percentageRevenueStr),
			ProfitInterest:     string(profitInterestStr),
			PerDay:             string(perDayStr),
			PerMounth:          string(perMounthStr),
			ConditionJobs:      string(conditionJobsStr),
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))
}

func (mc *MyClient) selectAccountDataId(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var data DataOwner
	r.ParseForm()
	idGet := string(r.Form.Get("id"))
	id, err := primitive.ObjectIDFromHex(strings.Trim(idGet, "\""))
	if err != nil {
		fmt.Println(err)
	}

	podcastsCollection := mc.db.Collection("owners")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.FindOne(
		ctx,
		bson.M{"_id": id}).Decode(&data)

	/* parsedData = DataOwner{
		Name:               string(data.Name),
		Phone:              string(data.phone),
		ContactPerson:      string(data.contactPerson),
		Proceedings:        string(data.proceedings),
		GroundsForContract: string(data.groundsForContract),
		AdditionContract:   string(data.additionContract),
		PercentageRevenue:  string(data.percentageRevenue),
		ProfitInterest:     string(data.profitInterest),
		PerDay:             string(data.perDay),
		PerMounth:          string(data.perMounth),
		ConditionJobs:      string(data.conditionJobs),
	} */

	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)

	w.Write([]byte(bytes))
}

func (mc *MyClient) deleteOwnerData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var data DataOwner
	var err error
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data.Id)
	podcastsCollection := mc.db.Collection("owners")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}
