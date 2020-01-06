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

type DataTechnicalService struct {
	auto                          string `json:"auto"`
	typeJob                       string `json:"typeJob"`
	autoRun                       string `json:"autoRun"`
	listJobs                      string `json:"listJobs"`
	contragent                    string `json:"contragent"`
	statestatePassengerSeat       string `json:"statestatePassengerSeat"`
	resultDyagnostic              string `json:"resultDyagnostic"`
	coastSparePart                string `json:"coastSparePart "`
	coastJobs                     string `json:"coastJobs"`
	statusRes                     string `json:"statusRes"`
	tyreBrand                     string `json:"tyreBrand"`
	bodyCabineDamage              string `json:"bodyCabineDamage"`
	autoCleanliness               string `json:"autoCleanliness"`
	overallInteriorCleanliness    string `json:"overallInteriorCleanliness"`
	stateCeling                   string `json:"stateCeling"`
	statePassengerSeat            string `json:"statePassengerSeat"`
	stateDriverSeat               string `json:"stateDriverSeat"`
	stateSeatbelt                 string `json:"stateSeatbelt"`
	stateSteeringWheelAndSwitches string `json:"stateSteeringWheelAndSwitches"`
	statePanel                    string `json:"statePanel"`
	stateSwitchKPP                string `json:"stateSwitchKPP"`
	windscreenCondition           string `json:"windscreenCondition"`
	stateLeftwindscreen           string `json:"stateLeftwindscreen"`
	trunkCondition                string `json:"trunkCondition"`
	stateTyre                     string `json:"stateTyre"`
	foreginLicenceRegistration    bool   `json:"stateTyre"`
	datePicker                    string `json:"datePicker"`
}

func (mc *MyClient) insertTechnicalServiceData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data DataTechnicalService
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	podcastsCollection := mc.db.Collection("technicalService")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"auto", data.auto},
		{"typeJob", data.typeJob},
		{"autoRun", data.autoRun},
		{"listJobs", data.listJobs},
		{"contragent", data.contragent},
		{"statestatePassengerSeat", data.statestatePassengerSeat},
		{"resultDyagnostic", data.resultDyagnostic},
		{"coastSparePart", data.coastSparePart},
		{"coastJobs", data.coastJobs},
		{"statusRes", data.statusRes},
		{"tyreBrand", data.tyreBrand},
		{"bodyCabineDamage", data.bodyCabineDamage},
		{"autoCleanliness", data.autoCleanliness},
		{"overallInteriorCleanliness", data.overallInteriorCleanliness},
		{"stateCeling", data.stateCeling},
		{"statePassengerSeat", data.statePassengerSeat},
		{"stateDriverSeat", data.stateDriverSeat},
		{"stateSeatbelt", data.stateSeatbelt},
		{"stateSteeringWheelAndSwitches", data.stateSteeringWheelAndSwitches},
		{"statePanel", data.statePanel},
		{"stateSwitchKPP", data.stateSwitchKPP},
		{"windscreenCondition", data.windscreenCondition},
		{"stateLeftwindscreen", data.stateLeftwindscreen},
		{"trunkCondition", data.trunkCondition},
		{"stateTyre", data.stateTyre},
		{"foreginLicenceRegistration", data.stateTyre},
		{"datePicker", data.datePicker},
		{"dateInsert", time.Now()},
		{"dateUpdate", nil},
	})
	if err != nil {
		log.Fatal(err)
	}
}
