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

type Automobile struct {
	Id                                string `json:"_id"`
	StatusRes                         string `json:"statusRes"`
	Picker                            string `json:"picker"`
	Brand                             string `json:"brand"`
	Model                             string `json:"model"`
	Owner                             string `json:"owner"`
	Category                          string `json:"category"`
	AutoRun                           string `json:"autoRun"`
	NameInsuranceCompany              string `json:"nameInsuranceCompany"`
	NumberInsuranceCompany            string `json:"numberInsuranceCompany"`
	OilChangeMileageKm                string `json:"oilChangeMileageKm"`
	TyreType                          string `json:"tyreType"`
	BrandTyre                         string `json:"brandTyre"`
	BeaconNumber                      string `json:"beaconNumber"`
	IMEIbeacon                        string `json:"IMEIbeacon"`
	AdditionalEquipment               string `json:"additionalEquipment"`
	SeriaAndNumberOfPTS               string `json:"seriaAndNumberOfPTS"`
	VIN                               string `json:"VIN"`
	NameTypeTS                        string `json:"nameTypeTS"`
	CategoryTS                        string `json:"categoryTS"`
	ChassisFrame                      string `json:"chassisFrame"`
	ModelNumberMotor                  string `json:"modelNumberMotor"`
	YearIssued                        string `json:"yearIssued"`
	ColorCabina                       string `json:"colorCabina"`
	EnginePower                       string `json:"enginePower"`
	EngineWorkingVolume               string `json:"engineWorkingVolume"`
	MotorType                         string `json:"motorType"`
	EcologyClaas                      string `json:"ecologyClaas"`
	AllwedMaxWeight                   string `json:"allwedMaxWeight"`
	WeightWithoutLoads                string `json:"weightWithoutLoads"`
	ForeginLicenceRegistration        bool   `json:"foreginLicenceRegistration"`
	NumberSymbol                      string `json:"numberSymbol"`
	WhoIssuedPTS                      string `json:"whoIssuedPTS"`
	SeriaAndNumberSTS                 string `json:"seriaAndNumberSTS"`
	StartOperationMenu                string `json:"startOperationMenu"`
	StartOperationDate                string `json:"startOperationDate"`
	FinishOperationDate               string `json:"finishOperationDate"`
	PeriodInsurancePolicyValidityDate string `json:"periodInsurancePolicyValidityDate"`
	TermValidityTODate                string `json:"termValidityTODate"`
	DateIssuedSTDate                  string `json:"dateIssuedSTDate"`
	Free                              string `json:"free"`
	Token                             string `json:"token"`
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
		{"statusRes", data.StatusRes},
		{"picker", data.Picker},
		{"brand", data.Brand},
		{"model", data.Model},
		{"token", data.Token},
		{"owner", data.Owner},
		{"category", data.Category},
		{"autoRun", data.AutoRun},
		{"nameInsuranceCompany", data.NameInsuranceCompany},
		{"numberInsuranceCompany", data.NumberInsuranceCompany},
		{"oilChangeMileageKm", data.OilChangeMileageKm},
		{"tyreType", data.TyreType},
		{"brandTyre", data.BrandTyre},
		{"beaconNumber", data.BeaconNumber},
		{"IMEIbeacon", data.IMEIbeacon},
		{"additionalEquipment", data.AdditionalEquipment},
		{"seriaAndNumberOfPTS", data.SeriaAndNumberOfPTS},
		{"VIN", data.VIN},
		{"nameTypeTS", data.NameTypeTS},
		{"categoryTS", data.CategoryTS},
		{"chassisFrame", data.ChassisFrame},
		{"modelNumberMotor", data.ModelNumberMotor},
		{"yearIssued", data.YearIssued},
		{"colorCabina", data.ColorCabina},
		{"enginePower", data.EnginePower},
		{"engineWorkingVolume", data.EngineWorkingVolume},
		{"motorType", data.MotorType},
		{"ecologyClaas", data.EcologyClaas},
		{"allwedMaxWeight", data.AllwedMaxWeight},
		{"weightWithoutLoads", data.WeightWithoutLoads},
		{"foreginLicenceRegistration", data.ForeginLicenceRegistration},
		{"numberSymbol", data.NumberSymbol},
		{"whoIssuedPTS", data.WhoIssuedPTS},
		{"seriaAndNumberSTS", data.SeriaAndNumberSTS},
		{"startOperationMenu", data.StartOperationMenu},
		{"startOperationDate", data.StartOperationDate},
		{"finishOperationDate", data.FinishOperationDate},
		{"periodInsurancePolicyValidityDate", data.PeriodInsurancePolicyValidityDate},
		{"termValidityTODate", data.TermValidityTODate},
		{"DateIssuedSTDate", data.DateIssuedSTDate},
		{"free", "1"},
		{"dateInsert", time.Now()},
		{"dateUpdate", nil},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateAutomobileData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data Automobile
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
	podcastsCollection := mc.db.Collection("automobiles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resultUpdate, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"statusRes":                         data.StatusRes,
				"picker":                            data.Picker,
				"brand":                             data.Brand,
				"model":                             data.Model,
				"owner":                             data.Owner,
				"category":                          data.Category,
				"autoRun":                           data.AutoRun,
				"nameInsuranceCompany":              data.NameInsuranceCompany,
				"numberInsuranceCompany":            data.NumberInsuranceCompany,
				"oilChangeMileageKm":                data.OilChangeMileageKm,
				"tyreType":                          data.TyreType,
				"brandTyre":                         data.BrandTyre,
				"beaconNumber":                      data.BeaconNumber,
				"IMEIbeacon":                        data.IMEIbeacon,
				"additionalEquipment":               data.AdditionalEquipment,
				"seriaAndNumberOfPTS":               data.SeriaAndNumberOfPTS,
				"VIN":                               data.VIN,
				"nameTypeTS":                        data.NameTypeTS,
				"categoryTS":                        data.CategoryTS,
				"chassisFrame":                      data.ChassisFrame,
				"modelNumberMotor":                  data.ModelNumberMotor,
				"yearIssued":                        data.YearIssued,
				"colorCabina":                       data.ColorCabina,
				"enginePower":                       data.EnginePower,
				"engineWorkingVolume":               data.EngineWorkingVolume,
				"motorType":                         data.MotorType,
				"ecologyClaas":                      data.EcologyClaas,
				"allwedMaxWeight":                   data.AllwedMaxWeight,
				"weightWithoutLoads":                data.WeightWithoutLoads,
				"foreginLicenceRegistration":        data.ForeginLicenceRegistration,
				"numberSymbol":                      data.NumberSymbol,
				"whoIssuedPTS":                      data.WhoIssuedPTS,
				"seriaAndNumberSTS":                 data.SeriaAndNumberSTS,
				"startOperationMenu":                data.StartOperationMenu,
				"startOperationDate":                data.StartOperationDate,
				"finishOperationDate":               data.FinishOperationDate,
				"periodInsurancePolicyValidityDate": data.PeriodInsurancePolicyValidityDate,
				"termValidityTODate":                data.TermValidityTODate,
				"DateIssuedSTDate":                  data.DateIssuedSTDate,
				"dateUpdate":                        time.Now(),
			},
		},
	)
	fmt.Println(resultUpdate.ModifiedCount) // output: 1
}

func (mc *MyClient) selectAutomobileData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("automobiles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.ParseForm()
	token := string(r.Form.Get("token"))
	cur, err := podcastsCollection.Find(ctx, bson.D{{"token", token}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData []Automobile
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		idJson, err := json.Marshal(result["_id"])
		statusResJson, err := json.Marshal(result["statusRes"])
		pickerJson, err := json.Marshal(result["picker"])
		brandJson, err := json.Marshal(result["brand"])
		modelJson, err := json.Marshal(result["model"])
		ownerJson, err := json.Marshal(result["owner"])
		categoryJson, err := json.Marshal(result["category"])
		autoRunJson, err := json.Marshal(result["autoRun"])
		nameInsuranceCompanyJson, err := json.Marshal(result["nameInsuranceCompany"])
		numberInsuranceCompanyJson, err := json.Marshal(result["numberInsuranceCompany"])
		oilChangeMileageKmJson, err := json.Marshal(result["oilChangeMileageKm"])
		tyreTypeJson, err := json.Marshal(result["tyreType"])
		brandTyreJson, err := json.Marshal(result["brandTyre"])
		beaconNumberJson, err := json.Marshal(result["beaconNumber"])
		IMEIbeaconJson, err := json.Marshal(result["IMEIbeacon"])
		additionalEquipmentJson, err := json.Marshal(result["additionalEquipment"])
		seriaAndNumberOfPTSJson, err := json.Marshal(result["seriaAndNumberOfPTS"])
		VINJson, err := json.Marshal(result["VIN"])
		nameTypeTSJson, err := json.Marshal(result["nameTypeTS"])
		categoryTSJson, err := json.Marshal(result["categoryTS"])
		chassisFrameJson, err := json.Marshal(result["chassisFrame"])
		modelNumberMotorJson, err := json.Marshal(result["modelNumberMotor"])
		yearIssuedJson, err := json.Marshal(result["yearIssued"])
		colorCabinaJson, err := json.Marshal(result["colorCabina"])
		enginePowerJson, err := json.Marshal(result["enginePower"])
		engineWorkingVolumeJson, err := json.Marshal(result["engineWorkingVolume"])
		motorTypeJson, err := json.Marshal(result["motorType"])
		ecologyClaasJson, err := json.Marshal(result["ecologyClaas"])
		allwedMaxWeightJson, err := json.Marshal(result["allwedMaxWeight"])
		weightWithoutLoadsJson, err := json.Marshal(result["weightWithoutLoads"])
		foreginLicenceRegistrationJson, err := json.Marshal(result["foreginLicenceRegistration"])
		numberSymbolJson, err := json.Marshal(result["numberSymbol"])
		whoIssuedPTSJson, err := json.Marshal(result["whoIssuedPTS"])
		seriaAndNumberSTSJson, err := json.Marshal(result["seriaAndNumberSTS"])
		startOperationMenuJson, err := json.Marshal(result["startOperationMenu"])
		startOperationDateJson, err := json.Marshal(result["startOperationDate"])
		finishOperationDateJson, err := json.Marshal(result["finishOperationDate"])
		periodInsurancePolicyValidityDateJson, err := json.Marshal(result["periodInsurancePolicyValidityDate"])
		termValidityTODateJson, err := json.Marshal(result["termValidityTODate"])
		dateIssuedSTSDateJson, err := json.Marshal(result["dateIssuedSTSDate"])

		idStr, _ := strconv.Unquote(string(idJson))
		statusResStr, _ := strconv.Unquote(string(statusResJson))
		pickerStr, _ := strconv.Unquote(string(pickerJson))
		brandStr, _ := strconv.Unquote(string(brandJson))
		modelStr, _ := strconv.Unquote(string(modelJson))
		ownerStr, _ := strconv.Unquote(string(ownerJson))
		categoryStr, _ := strconv.Unquote(string(categoryJson))
		autoRunStr, _ := strconv.Unquote(string(autoRunJson))
		nameInsuranceCompanyStr, _ := strconv.Unquote(string(nameInsuranceCompanyJson))
		numberInsuranceCompanyStr, _ := strconv.Unquote(string(numberInsuranceCompanyJson))
		oilChangeMileageKmStr, _ := strconv.Unquote(string(oilChangeMileageKmJson))
		tyreTypeStr, _ := strconv.Unquote(string(tyreTypeJson))
		brandTyreStr, _ := strconv.Unquote(string(brandTyreJson))
		beaconNumberStr, _ := strconv.Unquote(string(beaconNumberJson))
		IMEIbeaconStr, _ := strconv.Unquote(string(IMEIbeaconJson))
		additionalEquipmentStr, _ := strconv.Unquote(string(additionalEquipmentJson))
		seriaAndNumberOfPTSStr, _ := strconv.Unquote(string(seriaAndNumberOfPTSJson))
		VINStr, _ := strconv.Unquote(string(VINJson))
		nameTypeTSStr, _ := strconv.Unquote(string(nameTypeTSJson))
		categoryTSStr, _ := strconv.Unquote(string(categoryTSJson))
		chassisFrameStr, _ := strconv.Unquote(string(chassisFrameJson))
		modelNumberMotorStr, _ := strconv.Unquote(string(modelNumberMotorJson))
		yearIssuedStr, _ := strconv.Unquote(string(yearIssuedJson))
		colorCabinaStr, _ := strconv.Unquote(string(colorCabinaJson))
		enginePowerStr, _ := strconv.Unquote(string(enginePowerJson))
		engineWorkingVolumeStr, _ := strconv.Unquote(string(engineWorkingVolumeJson))
		motorTypeStr, _ := strconv.Unquote(string(motorTypeJson))
		ecologyClaasStr, _ := strconv.Unquote(string(ecologyClaasJson))
		allwedMaxWeightStr, _ := strconv.Unquote(string(allwedMaxWeightJson))
		weightWithoutLoadsStr, _ := strconv.Unquote(string(weightWithoutLoadsJson))
		foreginLicenceRegistrationStr, _ := strconv.ParseBool(string(foreginLicenceRegistrationJson))
		numberSymbolStr, _ := strconv.Unquote(string(numberSymbolJson))
		whoIssuedPTSStr, _ := strconv.Unquote(string(whoIssuedPTSJson))
		seriaAndNumberSTSStr, _ := strconv.Unquote(string(seriaAndNumberSTSJson))
		startOperationMenuStr, _ := strconv.Unquote(string(startOperationMenuJson))
		startOperationDateStr, _ := strconv.Unquote(string(startOperationDateJson))
		finishOperationDateStr, _ := strconv.Unquote(string(finishOperationDateJson))
		periodInsurancePolicyValidityDateStr, _ := strconv.Unquote(string(periodInsurancePolicyValidityDateJson))
		termValidityTODateStr, _ := strconv.Unquote(string(termValidityTODateJson))
		dateIssuedSTSDateStr, _ := strconv.Unquote(string(dateIssuedSTSDateJson))

		parsedData = append(parsedData, Automobile{
			Id:                                string(idStr),
			StatusRes:                         string(statusResStr),
			Picker:                            string(pickerStr),
			Brand:                             string(brandStr),
			Model:                             string(modelStr),
			Owner:                             string(ownerStr),
			Category:                          string(categoryStr),
			AutoRun:                           string(autoRunStr),
			NameInsuranceCompany:              string(nameInsuranceCompanyStr),
			NumberInsuranceCompany:            string(numberInsuranceCompanyStr),
			OilChangeMileageKm:                string(oilChangeMileageKmStr),
			TyreType:                          string(tyreTypeStr),
			BrandTyre:                         string(brandTyreStr),
			BeaconNumber:                      string(beaconNumberStr),
			IMEIbeacon:                        string(IMEIbeaconStr),
			AdditionalEquipment:               string(additionalEquipmentStr),
			SeriaAndNumberOfPTS:               string(seriaAndNumberOfPTSStr),
			VIN:                               string(VINStr),
			NameTypeTS:                        string(nameTypeTSStr),
			CategoryTS:                        string(categoryTSStr),
			ChassisFrame:                      string(chassisFrameStr),
			ModelNumberMotor:                  string(modelNumberMotorStr),
			YearIssued:                        string(yearIssuedStr),
			ColorCabina:                       string(colorCabinaStr),
			EnginePower:                       string(enginePowerStr),
			EngineWorkingVolume:               string(engineWorkingVolumeStr),
			MotorType:                         string(motorTypeStr),
			EcologyClaas:                      string(ecologyClaasStr),
			AllwedMaxWeight:                   string(allwedMaxWeightStr),
			WeightWithoutLoads:                string(weightWithoutLoadsStr),
			ForeginLicenceRegistration:        bool(foreginLicenceRegistrationStr),
			NumberSymbol:                      string(numberSymbolStr),
			WhoIssuedPTS:                      string(whoIssuedPTSStr),
			SeriaAndNumberSTS:                 string(seriaAndNumberSTSStr),
			StartOperationMenu:                string(startOperationMenuStr),
			StartOperationDate:                string(startOperationDateStr),
			FinishOperationDate:               string(finishOperationDateStr),
			PeriodInsurancePolicyValidityDate: string(periodInsurancePolicyValidityDateStr),
			TermValidityTODate:                string(termValidityTODateStr),
			DateIssuedSTDate:                  string(dateIssuedSTSDateStr),
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))
}

func (mc *MyClient) deleteAutomobileData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	//var data Automobile
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
	podcastsCollection := mc.db.Collection("automobiles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}

func (mc *MyClient) selectAutomobileDataOne(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var data Automobile
	//var parsedData Automobile
	r.ParseForm()
	idGet := string(r.Form.Get("id"))
	id, err := primitive.ObjectIDFromHex(strings.Trim(idGet, "\""))
	if err != nil {
		fmt.Println(err)
	}

	podcastsCollection := mc.db.Collection("automobiles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.FindOne(
		ctx,
		bson.M{"_id": id}).Decode(&data)

	data = Automobile{
		Id:                                string(data.Id),
		StatusRes:                         string(data.StatusRes),
		Picker:                            string(data.Picker),
		Brand:                             string(data.Brand),
		Model:                             string(data.Model),
		Owner:                             string(data.Owner),
		Category:                          string(data.Category),
		AutoRun:                           string(data.AutoRun),
		NameInsuranceCompany:              string(data.NameInsuranceCompany),
		NumberInsuranceCompany:            string(data.NumberInsuranceCompany),
		OilChangeMileageKm:                string(data.OilChangeMileageKm),
		TyreType:                          string(data.TyreType),
		BrandTyre:                         string(data.BrandTyre),
		BeaconNumber:                      string(data.BeaconNumber),
		IMEIbeacon:                        string(data.IMEIbeacon),
		AdditionalEquipment:               string(data.AdditionalEquipment),
		SeriaAndNumberOfPTS:               string(data.SeriaAndNumberOfPTS),
		VIN:                               string(data.VIN),
		NameTypeTS:                        string(data.NameTypeTS),
		CategoryTS:                        string(data.CategoryTS),
		ChassisFrame:                      string(data.ChassisFrame),
		ModelNumberMotor:                  string(data.ModelNumberMotor),
		YearIssued:                        string(data.YearIssued),
		ColorCabina:                       string(data.ColorCabina),
		EnginePower:                       string(data.EnginePower),
		EngineWorkingVolume:               string(data.EngineWorkingVolume),
		MotorType:                         string(data.MotorType),
		EcologyClaas:                      string(data.EcologyClaas),
		AllwedMaxWeight:                   string(data.AllwedMaxWeight),
		WeightWithoutLoads:                string(data.WeightWithoutLoads),
		ForeginLicenceRegistration:        bool(data.ForeginLicenceRegistration),
		NumberSymbol:                      string(data.NumberSymbol),
		WhoIssuedPTS:                      string(data.WhoIssuedPTS),
		SeriaAndNumberSTS:                 string(data.SeriaAndNumberSTS),
		StartOperationMenu:                string(data.StartOperationMenu),
		StartOperationDate:                string(data.StartOperationDate),
		FinishOperationDate:               string(data.FinishOperationDate),
		PeriodInsurancePolicyValidityDate: string(data.PeriodInsurancePolicyValidityDate),
		TermValidityTODate:                string(data.TermValidityTODate),
		DateIssuedSTDate:                  string(data.DateIssuedSTDate),
	}

	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)

	w.Write([]byte(bytes))
}
