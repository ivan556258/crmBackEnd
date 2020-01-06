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

type Automobile struct {
	statusRes                           string `json:"statusRes"`
	picker                              string `json:"picker"`
	valid                               bool   `json:"valid"`
	brand                               string `json:"brand"`
	model                               string `json:"model"`
	owner                               string `json:"owner"`
	category                            string `json:"category"`
	autoRun                             string `json:"autoRun"`
	nameInsuranceCompany                string `json:"nameInsuranceCompany"`
	numberInsuranceCompany              string `json:"numberInsuranceCompany"`
	oilChangeMileageKm                  string `json:"oilChangeMileageKm"`
	tyreType                            string `json:"tyreType"`
	brandTyre                           string `json:"brandTyre"`
	beaconNumber                        string `json:"beaconNumber"`
	IMEIbeacon                          string `json:"IMEIbeacon"`
	additionalEquipment                 string `json:"additionalEquipment"`
	seriaAndNumberOfPTS                 string `json:"seriaAndNumberOfPTS"`
	VIN                                 string `json:"VIN"`
	nameTypeTS                          string `json:"nameTypeTS"`
	categoryTS                          string `json:"categoryTS"`
	chassisFrame                        string `json:"chassisFrame"`
	modelNumberMotor                    string `json:"modelNumberMotor"`
	yearIssued                          string `json:"yearIssued"`
	colorCabina                         string `json:"colorCabina"`
	enginePower                         string `json:"enginePower"`
	engineWorkingVolume                 string `json:"engineWorkingVolume"`
	motorType                           string `json:"motorType"`
	ecologyClaas                        string `json:"ecologyClaas"`
	allwedMaxWeight                     string `json:"allwedMaxWeight"`
	weightWithoutLoads                  string `json:"weightWithoutLoads"`
	foreginLicenceRegistration          bool   `json:"foreginLicenceRegistration"`
	numberSymbol                        string `json:"numberSymbol"`
	whoIssuedPTS                        string `json:"whoIssuedPTS"`
	seriaAndNumberSTS                   string `json:"seriaAndNumberSTS"`
	startOperationMenu                  string `json:"startOperationMenu"`
	startOperationDate                  string `json:"startOperationDate"`
	startOperationPicker                string `json:"startOperationPicker"`
	finishOperationPicker               string `json:"finishOperationPicker"`
	periodInsurancePolicyValidityPicker string `json:"periodInsurancePolicyValidityPicker"`
	termValidityTOPicker                string `json:"termValidityTOPicker "`
	dateIssuedSTSPicker                 string `json:"dateIssuedSTSPicker"`
}

func (mc *MyClient) insertAutomobileData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data Automobile
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	podcastsCollection := mc.db.Collection("automobiles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"statusRes", data.statusRes},
		{"picker", data.picker},
		{"valid", data.valid},
		{"brand", data.brand},
		{"model", data.model},
		{"owner", data.owner},
		{"category", data.category},
		{"autoRun", data.autoRun},
		{"nameInsuranceCompany", data.nameInsuranceCompany},
		{"numberInsuranceCompany", data.numberInsuranceCompany},
		{"oilChangeMileageKm", data.oilChangeMileageKm},
		{"tyreType", data.tyreType},
		{"brandTyre", data.brandTyre},
		{"beaconNumber", data.beaconNumber},
		{"IMEIbeacon", data.IMEIbeacon},
		{"additionalEquipment", data.additionalEquipment},
		{"seriaAndNumberOfPTS", data.seriaAndNumberOfPTS},
		{"VIN", data.VIN},
		{"nameTypeTS", data.nameTypeTS},
		{"categoryTS", data.categoryTS},
		{"chassisFrame", data.chassisFrame},
		{"modelNumberMotor", data.modelNumberMotor},
		{"yearIssued", data.yearIssued},
		{"colorCabina", data.colorCabina},
		{"enginePower", data.enginePower},
		{"engineWorkingVolume", data.engineWorkingVolume},
		{"motorType", data.motorType},
		{"ecologyClaas", data.ecologyClaas},
		{"allwedMaxWeight", data.allwedMaxWeight},
		{"weightWithoutLoads", data.weightWithoutLoads},
		{"foreginLicenceRegistration", data.foreginLicenceRegistration},
		{"numberSymbol", data.numberSymbol},
		{"whoIssuedPTS", data.whoIssuedPTS},
		{"seriaAndNumberSTS", data.seriaAndNumberSTS},
		{"startOperationMenu", data.startOperationMenu},
		{"startOperationDate", data.startOperationDate},
		{"startOperationPicker", data.startOperationPicker},
		{"finishOperationPicker", data.finishOperationPicker},
		{"periodInsurancePolicyValidityPicker", data.periodInsurancePolicyValidityPicker},
		{"termValidityTOPicker", data.termValidityTOPicker},
		{"dateIssuedSTSPicker", data.dateIssuedSTSPicker},
		{"dateInsert", time.Now()},
		{"dateUpdate", nil},
	})
	if err != nil {
		log.Fatal(err)
	}
}
