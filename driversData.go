package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type DataDriver struct {
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
