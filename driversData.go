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

type DataDriver struct {
	Id                            string `json:"_id"`
	Lastname                      string `json:"lastname"`
	Firstname                     string `json:"firstname"`
	Fathername                    string `json:"fathername"`
	SeriaAndNumberPassport        string `json:"seriaAndNumberPassport"`
	LocationBrithday              string `json:"locationBrithday"`
	AddressRegistration           string `json:"addressRegistration"`
	IsOwner                       bool   `json:"isOwner"`
	Phone                         string `json:"phone"`
	Email                         string `json:"email"`
	Inn                           string `json:"inn"`
	ClassInsurance                string `json:"classInsurance"`
	NumberDriverLicence           string `json:"numberDriverLicence"`
	DateIssuedDriverLicenceDate   string `json:"dateIssuedDriverLicenceDate"`
	AddressInLifes                string `json:"addressInLifes"`
	MoreContacts                  string `json:"moreContacts"`
	ForeginDriversLicence         bool   `json:""`
	IsSelfCar                     bool   `json:"isSelfCar"`
	CarBrandAndNumber             string `json:"carBrandAndNumber"`
	Rating                        string `json:"rating"`
	Commentaries                  string `json:"commentaries"`
	InformDriverBalanceChanges    bool   `json:"informDriverBalanceChanges"`
	InformDriverBalanceLittle     bool   `json:"informDriverBalanceLittle"`
	InformDriverNewPenalty        bool   `json:"informDriverNewPenalty"`
	InformDriverOilChange         bool   `json:"informDriverOilChange"`
	AllowedBlocked                bool   `json:"allowedBlocked"`
	OnAutomaticRentMoney          bool   `json:"onAutomaticRentMoney"`
	ThresholdBalanceForDriver     string `json:"thresholdBalanceForDriver"`
	DateIssuedDate                string `json:"dateIssuedDate"`
	Brithday                      string `json:"brithday"`
	Issued                        string `json:"issued"`
	CodePollicia                  string `json:"codePollicia"`
	Brithdaypicker                string `json:"brithdaypicker"`
	DateIssuedPicker              string `json:"dateIssuedPicker"`
	DateIssuedDriverLicencePicker string `json:"dateIssuedDriverLicencePicker"`
	Status                        string `json:"status"`
	Token                         string `json:"token"`
}

func (mc *MyClient) insertDriverData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data DataDriver
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	podcastsCollection := mc.db.Collection("drivers")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"lastname", data.Lastname},
		{"firstname", data.Firstname},
		{"fathername", data.Fathername},
		{"seriaAndNumberPassport", data.SeriaAndNumberPassport},
		{"locationBrithday", data.LocationBrithday},
		{"addressRegistration", data.AddressRegistration},
		{"isOwner", data.IsOwner},
		{"phone", data.Phone},
		{"email", data.Email},
		{"token", data.Token},
		{"inn", data.Inn},
		{"classInsurance", data.ClassInsurance},
		{"numberDriverLicence", data.NumberDriverLicence},
		{"dateIssuedDriverLicenceDate", data.DateIssuedDriverLicenceDate},
		{"addressInLifes", data.AddressInLifes},
		{"moreContacts", data.MoreContacts},
		{"foreginDriversLicence", data.ForeginDriversLicence},
		{"isSelfCar", data.IsSelfCar},
		{"carBrandAndNumber", data.CarBrandAndNumber},
		{"rating", data.Rating},
		{"commentaries", data.Commentaries},
		{"informDriverBalanceChanges", data.InformDriverBalanceChanges},
		{"informDriverBalanceLittle", data.InformDriverBalanceLittle},
		{"informDriverNewPenalty", data.InformDriverNewPenalty},
		{"informDriverOilChange", data.InformDriverOilChange},
		{"allowedBlocked", data.AllowedBlocked},
		{"onAutomaticRentMoney", data.OnAutomaticRentMoney},
		{"thresholdBalanceForDriver", data.ThresholdBalanceForDriver},
		{"dateIssuedDate", data.DateIssuedDate},
		{"brithday", data.Brithday},
		{"issued", data.Issued},
		{"free", 1},
		{"codePollicia", data.CodePollicia},
		{"brithdaypicker", data.Brithdaypicker},
		{"dateIssuedPicker", data.DateIssuedPicker},
		{"dateIssuedDriverLicencePicker", data.DateIssuedDriverLicencePicker},
		{"status", data.Status},
		{"dateInsert", time.Now()},
		{"dateUpdate", nil},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateDriverData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data DataDriver
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
	podcastsCollection := mc.db.Collection("drivers")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resultUpdate, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"lastname":                      data.Lastname,
				"firstname":                     data.Firstname,
				"fathername":                    data.Fathername,
				"seriaAndNumberPassport":        data.SeriaAndNumberPassport,
				"locationBrithday":              data.LocationBrithday,
				"addressRegistration":           data.AddressRegistration,
				"isOwner":                       data.IsOwner,
				"phone":                         data.Phone,
				"email":                         data.Email,
				"inn":                           data.Inn,
				"classInsurance":                data.ClassInsurance,
				"numberDriverLicence":           data.NumberDriverLicence,
				"dateIssuedDriverLicenceDate":   data.DateIssuedDriverLicenceDate,
				"addressInLifes":                data.AddressInLifes,
				"moreContacts":                  data.MoreContacts,
				"foreginDriversLicence":         data.ForeginDriversLicence,
				"isSelfCar":                     data.IsSelfCar,
				"carBrandAndNumber":             data.CarBrandAndNumber,
				"rating":                        data.Rating,
				"commentaries":                  data.Commentaries,
				"informDriverBalanceChanges":    data.InformDriverBalanceChanges,
				"informDriverBalanceLittle":     data.InformDriverBalanceLittle,
				"informDriverNewPenalty":        data.InformDriverNewPenalty,
				"informDriverOilChange":         data.InformDriverOilChange,
				"allowedBlocked":                data.AllowedBlocked,
				"onAutomaticRentMoney":          data.OnAutomaticRentMoney,
				"thresholdBalanceForDriver":     data.ThresholdBalanceForDriver,
				"dateIssuedDate":                data.DateIssuedDate,
				"brithday":                      data.Brithday,
				"issued":                        data.Issued,
				"codePollicia":                  data.CodePollicia,
				"brithdaypicker":                data.Brithdaypicker,
				"dateIssuedPicker":              data.DateIssuedPicker,
				"dateIssuedDriverLicencePicker": data.DateIssuedDriverLicencePicker,
				"status":                        data.Status,
				"dateUpdate":                    time.Now(),
			},
		},
	)
	fmt.Println(resultUpdate.ModifiedCount) // output: 1
}

func (mc *MyClient) selectDriverData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("drivers")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.ParseForm()
	token := string(r.Form.Get("token"))
	cur, err := podcastsCollection.Find(ctx, bson.D{{"token", token}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData []DataDriver
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		idJson, err := json.Marshal(result["_id"])
		lastnameJson, err := json.Marshal(result["lastname"])
		firstnameJson, err := json.Marshal(result["firstname"])
		fathernameJson, err := json.Marshal(result["fathername"])
		seriaAndNumberPassportJson, err := json.Marshal(result["seriaAndNumberPassport"])
		locationBrithdayJson, err := json.Marshal(result["locationBrithday"])
		addressRegistrationJson, err := json.Marshal(result["addressRegistration"])
		isOwnerJson, err := json.Marshal(result["isOwner"])
		phoneJson, err := json.Marshal(result["phone"])
		emailJson, err := json.Marshal(result["email"])
		innJson, err := json.Marshal(result["inn"])
		classInsuranceJson, err := json.Marshal(result["classInsurance"])
		numberDriverLicenceJson, err := json.Marshal(result["numberDriverLicence"])
		dateIssuedDriverLicenceDateJson, err := json.Marshal(result["dateIssuedDriverLicenceDate"])
		addressInLifesJson, err := json.Marshal(result["addressInLifes"])
		moreContactsJson, err := json.Marshal(result["moreContacts"])
		foreginDriversLicenceJson, err := json.Marshal(result["foreginDriversLicence"])
		isSelfCarJson, err := json.Marshal(result["isSelfCar"])
		carBrandAndNumberJson, err := json.Marshal(result["carBrandAndNumber"])
		ratingJson, err := json.Marshal(result["rating"])
		commentariesJson, err := json.Marshal(result["commentaries"])
		informDriverBalanceChangesJson, err := json.Marshal(result["informDriverBalanceChanges"])
		informDriverBalanceLittleJson, err := json.Marshal(result["informDriverBalanceLittle"])
		informDriverNewPenaltyJson, err := json.Marshal(result["informDriverNewPenalty"])
		informDriverOilChangeJson, err := json.Marshal(result["informDriverOilChange"])
		allowedBlockedJson, err := json.Marshal(result["allowedBlocked"])
		onAutomaticRentMoneyJson, err := json.Marshal(result["onAutomaticRentMoney"])
		thresholdBalanceForDriverJson, err := json.Marshal(result["thresholdBalanceForDriver"])
		dateIssuedDateJson, err := json.Marshal(result["dateIssuedDate"])
		brithdayJson, err := json.Marshal(result["brithday"])
		issuedJson, err := json.Marshal(result["issued"])
		codePolliciaJson, err := json.Marshal(result["codePollicia"])
		brithdaypickerJson, err := json.Marshal(result["brithdaypicker"])
		dateIssuedPickerJson, err := json.Marshal(result["dateIssuedPicker"])
		dateIssuedDriverLicencePickerJson, err := json.Marshal(result["dateIssuedDriverLicencePicker"])
		statusJson, err := json.Marshal(result["status"])

		idStr, _ := strconv.Unquote(string(idJson))
		lastnameStr, _ := strconv.Unquote(string(lastnameJson))
		firstnameStr, _ := strconv.Unquote(string(firstnameJson))
		fathernameStr, _ := strconv.Unquote(string(fathernameJson))
		seriaAndNumberPassportStr, _ := strconv.Unquote(string(seriaAndNumberPassportJson))
		locationBrithdayStr, _ := strconv.Unquote(string(locationBrithdayJson))
		addressRegistrationStr, _ := strconv.Unquote(string(addressRegistrationJson))
		isOwnerStr, _ := strconv.ParseBool(string(isOwnerJson))
		phoneStr, _ := strconv.Unquote(string(phoneJson))
		emailStr, _ := strconv.Unquote(string(emailJson))
		innStr, _ := strconv.Unquote(string(innJson))
		classInsuranceStr, _ := strconv.Unquote(string(classInsuranceJson))
		numberDriverLicenceStr, _ := strconv.Unquote(string(numberDriverLicenceJson))
		dateIssuedDriverLicenceDateStr, _ := strconv.Unquote(string(dateIssuedDriverLicenceDateJson))
		addressInLifesStr, _ := strconv.Unquote(string(addressInLifesJson))
		moreContactsStr, _ := strconv.Unquote(string(moreContactsJson))
		foreginDriversLicenceStr, _ := strconv.ParseBool(string(foreginDriversLicenceJson))
		isSelfCarStr, _ := strconv.ParseBool(string(isSelfCarJson))
		carBrandAndNumberStr, _ := strconv.Unquote(string(carBrandAndNumberJson))
		ratingStr, _ := strconv.Unquote(string(ratingJson))
		commentariesStr, _ := strconv.Unquote(string(commentariesJson))
		informDriverBalanceChangesStr, _ := strconv.ParseBool(string(informDriverBalanceChangesJson))
		informDriverBalanceLittleStr, _ := strconv.ParseBool(string(informDriverBalanceLittleJson))
		informDriverNewPenaltyStr, _ := strconv.ParseBool(string(informDriverNewPenaltyJson))
		informDriverOilChangeStr, _ := strconv.ParseBool(string(informDriverOilChangeJson))
		allowedBlockedStr, _ := strconv.ParseBool(string(allowedBlockedJson))
		onAutomaticRentMoneyStr, _ := strconv.ParseBool(string(onAutomaticRentMoneyJson))
		thresholdBalanceForDriverStr, _ := strconv.Unquote(string(thresholdBalanceForDriverJson))
		dateIssuedDateStr, _ := strconv.Unquote(string(dateIssuedDateJson))
		brithdayStr, _ := strconv.Unquote(string(brithdayJson))
		issuedStr, _ := strconv.Unquote(string(issuedJson))
		codePolliciaStr, _ := strconv.Unquote(string(codePolliciaJson))
		brithdaypickerStr, _ := strconv.Unquote(string(brithdaypickerJson))
		dateIssuedPickerStr, _ := strconv.Unquote(string(dateIssuedPickerJson))
		dateIssuedDriverLicencePickerStr, _ := strconv.Unquote(string(dateIssuedDriverLicencePickerJson))
		statusStr, _ := strconv.Unquote(string(statusJson))

		parsedData = append(parsedData, DataDriver{
			Id:                            string(idStr),
			Lastname:                      string(lastnameStr),
			Firstname:                     string(firstnameStr),
			Fathername:                    string(fathernameStr),
			SeriaAndNumberPassport:        string(seriaAndNumberPassportStr),
			LocationBrithday:              string(locationBrithdayStr),
			AddressRegistration:           string(addressRegistrationStr),
			IsOwner:                       bool(isOwnerStr),
			Phone:                         string(phoneStr),
			Email:                         string(emailStr),
			Inn:                           string(innStr),
			ClassInsurance:                string(classInsuranceStr),
			NumberDriverLicence:           string(numberDriverLicenceStr),
			DateIssuedDriverLicenceDate:   string(dateIssuedDriverLicenceDateStr),
			AddressInLifes:                string(addressInLifesStr),
			MoreContacts:                  string(moreContactsStr),
			ForeginDriversLicence:         bool(foreginDriversLicenceStr),
			IsSelfCar:                     bool(isSelfCarStr),
			CarBrandAndNumber:             string(carBrandAndNumberStr),
			Rating:                        string(ratingStr),
			Commentaries:                  string(commentariesStr),
			InformDriverBalanceChanges:    bool(informDriverBalanceChangesStr),
			InformDriverBalanceLittle:     bool(informDriverBalanceLittleStr),
			InformDriverNewPenalty:        bool(informDriverNewPenaltyStr),
			InformDriverOilChange:         bool(informDriverOilChangeStr),
			AllowedBlocked:                bool(allowedBlockedStr),
			OnAutomaticRentMoney:          bool(onAutomaticRentMoneyStr),
			ThresholdBalanceForDriver:     string(thresholdBalanceForDriverStr),
			DateIssuedDate:                string(dateIssuedDateStr),
			Brithday:                      string(brithdayStr),
			Issued:                        string(issuedStr),
			CodePollicia:                  string(codePolliciaStr),
			Brithdaypicker:                string(brithdaypickerStr),
			DateIssuedPicker:              string(dateIssuedPickerStr),
			DateIssuedDriverLicencePicker: string(dateIssuedDriverLicencePickerStr),
			Status:                        string(statusStr),
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))
}

func (mc *MyClient) deleteDriverData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data DataDriver
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
	podcastsCollection := mc.db.Collection("drivers")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}

func (mc *MyClient) selectDriverDataOne(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var data DataDriver
	r.ParseForm()
	idGet := string(r.Form.Get("id"))
	id, err := primitive.ObjectIDFromHex(strings.Trim(idGet, "\""))
	if err != nil {
		fmt.Println(err)
	}

	podcastsCollection := mc.db.Collection("drivers")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.FindOne(
		ctx,
		bson.M{"_id": id}).Decode(&data)

	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)

	w.Write([]byte(bytes))
}
