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

type UserCompany struct {
	Id                   string `json:"_id"`
	Name                 string `json:"name"`
	ContractTrack        string `json:"contractTrack"`
	AutoReturn           string `json:"autoReturn"`
	CustomToolbar        string `json:"customToolbar"`
	Congratulations      string `json:"congratulations"`
	APIYandexKey         string `json:"APIYandexKey"`
	GlobAllowedBlockAuto bool   `json:"globAllowedBlockAuto"`
	SmsDriverBlock       string `json:"smsDriverBlock"`
	SmsDriverUnblock     string `json:"smsDriverUnblock"`
	IdPark               string `json:"idPark"`
	IdClient             string `json:"idClient"`
	ScoreYandexAPI       string `json:"scoreYandexAPI"`
	City                 string `json:"city"`
	Address              string `json:"address"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	PhoneSTO             string `json:"phoneSTO"`
	OnTurnSMS            bool   `json:"onTurnSMS"`
	SendHello            bool   `json:"sendHello"`
	TempleteSMSOne       string `json:"templeteSMSOne"`
	InformAboutAddScore  bool   `json:"informAboutAddScore"`
	TempleteSMSTwo       string `json:"templeteSMSTwo"`
	InformNewPenalty     bool   `json:"informNewPenalty"`
	TempleteSMSThree     string `json:"templeteSMSThree"`
	InformNeedChangeOli  bool   `json:"informNeedChangeOli"`
	TempleteSMSFoo       string `json:"templeteSMSFoo"`
	TempleteSMSFive      string `json:"templeteSMSFive"`
}

var data UserCompany
var parsedData UserCompany

func (mc *MyClient) insertUserCompanyData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	podcastsCollection := mc.db.Collection("userCompany")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"name", data.Name},
		{"contractTrack", data.ContractTrack},
		{"autoReturn", data.AutoReturn},
		{"customToolbar", data.CustomToolbar},
		{"congratulations", data.Congratulations},
		{"APIYandexKey", data.APIYandexKey},
		{"globAllowedBlockAuto", data.GlobAllowedBlockAuto},
		{"smsDriverBlock", data.SmsDriverBlock},
		{"smsDriverUnblock", data.SmsDriverUnblock},
		{"idPark", data.IdPark},
		{"idClient", data.IdClient},
		{"scoreYandexAPI", data.ScoreYandexAPI},
		{"city", data.City},
		{"address", data.Address},
		{"email", data.Email},
		{"phone", data.Phone},
		{"phoneSTO", data.PhoneSTO},
		{"onTurnSMS", data.OnTurnSMS},
		{"sendHello", data.SendHello},
		{"templeteSMSOne", data.TempleteSMSOne},
		{"informAboutAddScore", data.InformAboutAddScore},
		{"templeteSMSTwo", data.TempleteSMSTwo},
		{"informNewPenalty", data.InformNewPenalty},
		{"templeteSMSThree", data.TempleteSMSThree},
		{"informNeedChangeOli", data.InformNeedChangeOli},
		{"templeteSMSFoo", data.TempleteSMSFoo},
		{"templeteSMSFive", data.TempleteSMSFive},
		{"dateInsert", time.Now()},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) saveUserCompanyData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
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
	podcastsCollection := mc.db.Collection("userCompany")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resultUpdate, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"name":                 data.Name,
				"contractTrack":        data.ContractTrack,
				"autoReturn":           data.AutoReturn,
				"customToolbar":        data.CustomToolbar,
				"congratulations":      data.Congratulations,
				"APIYandexKey":         data.APIYandexKey,
				"globAllowedBlockAuto": data.GlobAllowedBlockAuto,
				"smsDriverBlock":       data.SmsDriverBlock,
				"smsDriverUnblock":     data.SmsDriverUnblock,
				"idPark":               data.IdPark,
				"idClient":             data.IdClient,
				"scoreYandexAPI":       data.ScoreYandexAPI,
				"city":                 data.City,
				"address":              data.Address,
				"email":                data.Email,
				"phone":                data.Phone,
				"phoneSTO":             data.PhoneSTO,
				"onTurnSMS":            data.OnTurnSMS,
				"sendHello":            data.SendHello,
				"templeteSMSOne":       data.TempleteSMSOne,
				"informAboutAddScore":  data.InformAboutAddScore,
				"templeteSMSTwo":       data.TempleteSMSTwo,
				"informNewPenalty":     data.InformNewPenalty,
				"templeteSMSThree":     data.TempleteSMSThree,
				"informNeedChangeOli":  data.InformNeedChangeOli,
				"templeteSMSFoo":       data.TempleteSMSFoo,
				"templeteSMSFive":      data.TempleteSMSFive,
				"dateUpdate":           time.Now(),
			},
		},
	)
	fmt.Println(resultUpdate.ModifiedCount) // output: 1
}

func (mc *MyClient) selectUserCompanyDataOne(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("userCompany")
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

func (mc *MyClient) deleteUserCompanyDataOne(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
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
	podcastsCollection := mc.db.Collection("account")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}
