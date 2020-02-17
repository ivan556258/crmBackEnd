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

type Data struct {
	Id          string `json:"_id"`
	Number      string `json:"number"`
	Driver      string `json:"driver"`
	Auto        string `json:"auto"`
	Tariff      string `json:"tariff"`
	DriverStr   string `json:"driverStr"`
	AutoStr     string `json:"autoStr"`
	TariffStr   string `json:"tariffStr"`
	Begindate   string `json:"begindate"`
	Enddate     string `json:"enddate"`
	Continues   bool   `json:"continues"`
	MoreInfo    string `json:"moreInfo"`
	Status      string `json:"status"`
	Token       string `json:"token"`
	DriverPhone string `json:"driverPhone"`
}

func (mc *MyClient) insertContractData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data Data
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	var driver string = mc.dataNameFind("drivers", "_id", data.Driver)
	var tariff string = mc.dataAccountBillItemFind("accountBillItem", "_id", data.Tariff)
	var auto string = mc.dataAutoFind("automobiles", "_id", data.Auto)

	podcastsCollection := mc.db.Collection("contractsRent")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"number", data.Number},
		{"driver", data.Driver},
		{"auto", data.Auto},
		{"tariff", data.Tariff},
		{"driverStr", driver},
		{"autoStr", auto},

		{"tariffStr", tariff},
		{"begindate", data.Begindate},
		{"enddate", data.Enddate},
		{"continues", data.Continues},
		{"moreInfo", data.MoreInfo},
		{"token", data.Token},
		{"status", data.Status},
		{"dateInsert", time.Now()},
		{"dateUpdate", nil},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateContractData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data Data
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

	var driver string = mc.dataNameFind("drivers", "_id", data.Driver)
	var tariff string = mc.dataAccountBillItemFind("accountBillItem", "_id", data.Tariff)
	var auto string = mc.dataAutoFind("automobiles", "_id", data.Auto)
	podcastsCollection := mc.db.Collection("contractsRent")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resultUpdate, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"number":     data.Number,
				"driver":     data.Driver,
				"auto":       data.Auto,
				"tariff":     data.Tariff,
				"driverStr":  driver,
				"autoStr":    auto,
				"tariffStr":  tariff,
				"begindate":  data.Begindate,
				"enddate":    data.Enddate,
				"continues":  data.Continues,
				"moreInfo":   data.MoreInfo,
				"status":     data.Status,
				"dateUpdate": time.Now(),
			},
		},
	)
	fmt.Println(resultUpdate.ModifiedCount) // output: 1
}

func (mc *MyClient) selectContractData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("contractsRent")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.ParseForm()
	token := string(r.Form.Get("token"))
	cur, err := podcastsCollection.Find(ctx, bson.D{{"token", token}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData []Data
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		idJson, err := json.Marshal(result["_id"])
		numberJson, err := json.Marshal(result["number"])
		driverJson, err := json.Marshal(result["driver"])
		autoJson, err := json.Marshal(result["auto"])
		tariffJson, err := json.Marshal(result["tariff"])
		driverStrJson, err := json.Marshal(result["driverStr"])
		autoStrJson, err := json.Marshal(result["autoStr"])
		tariffStrJson, err := json.Marshal(result["tariffStr"])
		begindateJson, err := json.Marshal(result["begindate"])
		enddateJson, err := json.Marshal(result["enddate"])
		continuesJson, err := json.Marshal(result["continues"])
		moreInfoJson, err := json.Marshal(result["moreInfo"])
		statusJson, err := json.Marshal(result["status"])

		idStr, _ := strconv.Unquote(string(idJson))
		numberStr, _ := strconv.Unquote(string(numberJson))
		driverStr, _ := strconv.Unquote(string(driverJson))
		autoStr, _ := strconv.Unquote(string(autoJson))
		tariffStr, _ := strconv.Unquote(string(tariffJson))
		driverStrStr, _ := strconv.Unquote(string(driverStrJson))
		autoStrStr, _ := strconv.Unquote(string(autoStrJson))
		tariffStrStr, _ := strconv.Unquote(string(tariffStrJson))
		begindateStr, _ := strconv.Unquote(string(begindateJson))
		enddateStr, _ := strconv.Unquote(string(enddateJson))
		continuesStr, _ := strconv.ParseBool(string(continuesJson))
		moreInfoStr, _ := strconv.Unquote(string(moreInfoJson))
		statusStr, _ := strconv.Unquote(string(statusJson))

		dataDriver := mc.dataDriverFind("drivers", "_id", string(driverStr))

		parsedData = append(parsedData, Data{
			Id:          string(idStr),
			Number:      string(numberStr),
			Driver:      string(driverStr),
			Auto:        string(autoStr),
			Tariff:      string(tariffStr),
			DriverStr:   string(driverStrStr),
			AutoStr:     string(autoStrStr),
			TariffStr:   string(tariffStrStr),
			Begindate:   string(begindateStr),
			Enddate:     string(enddateStr),
			Continues:   bool(continuesStr),
			MoreInfo:    string(moreInfoStr),
			Status:      string(statusStr),
			DriverPhone: string(dataDriver.Phone),
		})

	}
	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(parsedData)
	w.Write([]byte(bytes))
}

func moreInfoDriver() {

}

func (mc *MyClient) deleteContractData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data Data
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
	podcastsCollection := mc.db.Collection("contractsRent")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}

func (mc *MyClient) selectContractDataOne(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("contractsRent")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := podcastsCollection.Find(ctx, bson.D{})
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
		templeteSMSOneJson, err := json.Marshal(result["templeteSMSOne"])
		globAllowedBlockAutoJson, err := json.Marshal(result["globAllowedBlockAuto"])
		contractTrackJson, err := json.Marshal(result["contractTrack"])
		autoReturnJson, err := json.Marshal(result["autoReturn"])
		customToolbarJson, err := json.Marshal(result["customToolbar"])
		congratulationsJson, err := json.Marshal(result["congratulations"])
		APIYandexKeyJson, err := json.Marshal(result["APIYandexKey"])
		smsDriverBlockJson, err := json.Marshal(result["smsDriverBlock"])
		smsDriverUnblockJson, err := json.Marshal(result["smsDriverUnblock"])
		idParkJson, err := json.Marshal(result["idPark"])
		idClientJson, err := json.Marshal(result["idClient"])
		scoreYandexAPIJson, err := json.Marshal(result["scoreYandexAPI"])
		cityJson, err := json.Marshal(result["city"])
		addressJson, err := json.Marshal(result["address"])
		emailJson, err := json.Marshal(result["email"])
		phoneJson, err := json.Marshal(result["phone"])
		phoneSTOJson, err := json.Marshal(result["phoneSTO"])
		onTurnSMSJson, err := json.Marshal(result["onTurnSMS"])
		sendHelloJson, err := json.Marshal(result["sendHello"])
		informAboutAddScoreJson, err := json.Marshal(result["informAboutAddScore"])
		templeteSMSTwoJson, err := json.Marshal(result["templeteSMSTwo"])
		informNewPenaltyJson, err := json.Marshal(result["informNewPenalty"])
		templeteSMSThreeJson, err := json.Marshal(result["templeteSMSThree"])
		informNeedChangeOliJson, err := json.Marshal(result["informNeedChangeOli"])
		templeteSMSFooJson, err := json.Marshal(result["templeteSMSFoo"])
		templeteSMSFiveJson, err := json.Marshal(result["templeteSMSFive"])

		sid, _ := strconv.Unquote(string(id))
		nameStr, _ := strconv.Unquote(string(nameJson))
		templeteSMSOneStr, _ := strconv.Unquote(string(templeteSMSOneJson))
		contractTrackStr, _ := strconv.Unquote(string(contractTrackJson))
		autoReturnStr, _ := strconv.Unquote(string(autoReturnJson))
		customToolbarStr, _ := strconv.Unquote(string(customToolbarJson))
		congratulationsStr, _ := strconv.Unquote(string(congratulationsJson))
		APIYandexKeyStr, _ := strconv.Unquote(string(APIYandexKeyJson))
		smsDriverBlockStr, _ := strconv.Unquote(string(smsDriverBlockJson))
		smsDriverUnblockStr, _ := strconv.Unquote(string(smsDriverUnblockJson))
		idParkStr, _ := strconv.Unquote(string(idParkJson))
		idClientStr, _ := strconv.Unquote(string(idClientJson))
		scoreYandexAPIStr, _ := strconv.Unquote(string(scoreYandexAPIJson))
		cityStr, _ := strconv.Unquote(string(cityJson))
		addressStr, _ := strconv.Unquote(string(addressJson))
		emailStr, _ := strconv.Unquote(string(emailJson))
		phoneStr, _ := strconv.Unquote(string(phoneJson))
		phoneSTOStr, _ := strconv.Unquote(string(phoneSTOJson))
		templeteSMSTwoStr, _ := strconv.Unquote(string(templeteSMSTwoJson))
		templeteSMSThreeStr, _ := strconv.Unquote(string(templeteSMSThreeJson))
		templeteSMSFooStr, _ := strconv.Unquote(string(templeteSMSFooJson))
		templeteSMSFiveStr, _ := strconv.Unquote(string(templeteSMSFiveJson))

		onTurnSMSBool, _ := strconv.ParseBool(string(onTurnSMSJson))
		sendHelloBool, _ := strconv.ParseBool(string(sendHelloJson))
		informNewPenaltyBool, _ := strconv.ParseBool(string(informNewPenaltyJson))
		informAboutAddScoreBool, _ := strconv.ParseBool(string(informAboutAddScoreJson))
		globAllowedBlockAutoBool, _ := strconv.ParseBool(string(globAllowedBlockAutoJson))
		informNeedChangeOliBool, _ := strconv.ParseBool(string(informNeedChangeOliJson))

		parsedData = UserCompany{
			Id:                   string(sid),
			Name:                 string(nameStr),
			ContractTrack:        string(contractTrackStr),
			AutoReturn:           string(autoReturnStr),
			CustomToolbar:        string(customToolbarStr),
			Congratulations:      string(congratulationsStr),
			APIYandexKey:         string(APIYandexKeyStr),
			GlobAllowedBlockAuto: bool(globAllowedBlockAutoBool),
			SmsDriverBlock:       string(smsDriverBlockStr),
			SmsDriverUnblock:     string(smsDriverUnblockStr),
			IdPark:               string(idParkStr),
			IdClient:             string(idClientStr),
			ScoreYandexAPI:       string(scoreYandexAPIStr),
			City:                 string(cityStr),
			Address:              string(addressStr),
			Email:                string(emailStr),
			Phone:                string(phoneStr),
			PhoneSTO:             string(phoneSTOStr),
			OnTurnSMS:            bool(onTurnSMSBool),
			SendHello:            bool(sendHelloBool),
			TempleteSMSOne:       string(templeteSMSOneStr),
			InformAboutAddScore:  bool(informAboutAddScoreBool),
			TempleteSMSTwo:       string(templeteSMSTwoStr),
			InformNewPenalty:     bool(informNewPenaltyBool),
			TempleteSMSThree:     string(templeteSMSThreeStr),
			InformNeedChangeOli:  bool(informNeedChangeOliBool),
			TempleteSMSFoo:       string(templeteSMSFooStr),
			TempleteSMSFive:      string(templeteSMSFiveStr),
		}

	}
	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(parsedData)
	w.Write([]byte(bytes))
}
