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

type DataTechnicalService struct {
	Id                            string      `json:"_id"`
	Auto                          string      `json:"auto"`
	TypeJob                       string      `json:"typeJob"`
	AutoRun                       string      `json:"autoRun"`
	Contragent                    string      `json:"contragent"`
	StatestatePassengerSeat       string      `json:"statestatePassengerSeat"`
	ResultDyagnostic              string      `json:"resultDyagnostic"`
	CoastSparePart                string      `json:"coastSparePart"`
	CoastJobs                     interface{} `json:"coastJobs"`
	StatusRes                     string      `json:"statusRes"`
	TyreBrand                     string      `json:"tyreBrand"`
	BodyCabineDamage              string      `json:"bodyCabineDamage"`
	AutoCleanliness               string      `json:"autoCleanliness"`
	OverallInteriorCleanliness    string      `json:"overallInteriorCleanliness"`
	StateCeling                   string      `json:"stateCeling"`
	StatePassengerSeat            string      `json:"statePassengerSeat"`
	StateDriverSeat               string      `json:"stateDriverSeat"`
	StateSeatbelt                 string      `json:"stateSeatbelt"`
	StateSteeringWheelAndSwitches string      `json:"stateSteeringWheelAndSwitches"`
	StatePanel                    string      `json:"statePanel"`
	StateSwitchKPP                string      `json:"stateSwitchKPP"`
	WindscreenCondition           string      `json:"windscreenCondition"`
	StateLeftwindscreen           string      `json:"stateLeftwindscreen"`
	TrunkCondition                string      `json:"trunkCondition"`
	StateTyre                     string      `json:"stateTyre"`
	ForeginLicenceRegistration    bool        `json:"foreginLicenceRegistration"`
	DateData                      string      `json:"dateData"`
	Token                         string      `json:"token"`
	Parts                         interface{} `json:"parts"`
	Free                          string      `json:"free"`
	CoastJobsAll                  string      `json:"coastJobsAll"`
	OtherParts                    interface{} `json:"otherPart"`
}

type DataTechnicalServicex struct {
	Id                            string             `json:"_id"`
	Auto                          string             `json:"auto"`
	TypeJob                       string             `json:"typeJob"`
	AutoRun                       string             `json:"autoRun"`
	Contragent                    string             `json:"contragent"`
	StatestatePassengerSeat       string             `json:"statestatePassengerSeat"`
	ResultDyagnostic              string             `json:"resultDyagnostic"`
	CoastSparePart                string             `json:"coastSparePart"`
	CoastJobs                     []CoastJobsStruct  `json:"coastJobs"`
	StatusRes                     string             `json:"statusRes"`
	TyreBrand                     string             `json:"tyreBrand"`
	BodyCabineDamage              string             `json:"bodyCabineDamage"`
	AutoCleanliness               string             `json:"autoCleanliness"`
	OverallInteriorCleanliness    string             `json:"overallInteriorCleanliness"`
	StateCeling                   string             `json:"stateCeling"`
	StatePassengerSeat            string             `json:"statePassengerSeat"`
	StateDriverSeat               string             `json:"stateDriverSeat"`
	StateSeatbelt                 string             `json:"stateSeatbelt"`
	StateSteeringWheelAndSwitches string             `json:"stateSteeringWheelAndSwitches"`
	StatePanel                    string             `json:"statePanel"`
	StateSwitchKPP                string             `json:"stateSwitchKPP"`
	WindscreenCondition           string             `json:"windscreenCondition"`
	StateLeftwindscreen           string             `json:"stateLeftwindscreen"`
	TrunkCondition                string             `json:"trunkCondition"`
	StateTyre                     string             `json:"stateTyre"`
	ForeginLicenceRegistration    bool               `json:"foreginLicenceRegistration"`
	DateData                      string             `json:"dateData"`
	Token                         string             `json:"token"`
	Parts                         []PartsStruct      `json:"parts"`
	Free                          string             `json:"free"`
	CoastJobsAll                  string             `json:"coastJobsAll"`
	OtherParts                    []OtherPartsStruct `json:"otherPart"`
}

type CoastJobsStruct struct {
	Title   string `json:"title"`
	Price   string `json:"price"`
	Percent string `json:"percent"`
}

type PartsStruct struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Percent     string `json:"percent"`
	SummEnd     string `json:"summEnd"`
	Summ        string `json:"summ"`
}

type OtherPartsStruct struct {
	Name     string `json:"name"`
	Articale string `json:"articale"`
	HowMuch  string `json:"howmuch"`
	Price    string `json:"price"`
	Percent  string `json:"percent"`
	SummEnd  string `json:"summEnd"`
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

	strPartsStruct, _ := data.Parts.(string)
	strOtherPartsStruct, _ := data.OtherParts.(string)
	strCoastJobsStruct, _ := data.CoastJobs.(string)

	var i []PartsStruct
	var y []OtherPartsStruct
	var z []CoastJobsStruct

	var bMPartsStruct []bson.M
	var bMstrOtherPartsStruct []bson.M
	var bMstrCoastJobsStruct []bson.M

	if err := json.Unmarshal([]byte(strPartsStruct), &i); err != nil {
		fmt.Println("ugh: ", err)
	}

	if err := json.Unmarshal([]byte(strOtherPartsStruct), &y); err != nil {
		fmt.Println("ugh: ", err)
	}

	if err := json.Unmarshal([]byte(strCoastJobsStruct), &z); err != nil {
		fmt.Println("ugh: ", err)
	}

	for _, u := range i {
		bMPartsStruct = append(bMPartsStruct, bson.M{
			"id":          u.Id,
			"title":       u.Title,
			"description": u.Description,
			"percent":     u.Percent,
			"summEnd":     u.SummEnd,
			"summ":        u.Summ,
		})
	}
	for _, u := range y {
		bMstrOtherPartsStruct = append(bMstrOtherPartsStruct, bson.M{
			"name":     u.Name,
			"articale": u.Articale,
			"howMuch":  u.HowMuch,
			"price":    u.Price,
			"percent":  u.Percent,
			"summEnd":  u.SummEnd,
		})
	}
	for _, u := range z {
		bMstrCoastJobsStruct = append(bMstrCoastJobsStruct, bson.M{
			"title":   u.Title,
			"price":   u.Price,
			"percent": u.Percent,
		})
	}

	podcastsCollection := mc.db.Collection("technicalService")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.M{
		"auto":                          data.Auto,
		"typeJob":                       data.TypeJob,
		"autoRun":                       data.AutoRun,
		"coastJobsAll":                  data.CoastJobsAll,
		"contragent":                    data.Contragent,
		"statestatePassengerSeat":       data.StatestatePassengerSeat,
		"resultDyagnostic":              data.ResultDyagnostic,
		"coastSparePart":                data.CoastSparePart,
		"coastJobs":                     bMstrCoastJobsStruct,
		"statusRes":                     data.StatusRes,
		"tyreBrand":                     data.TyreBrand,
		"bodyCabineDamage":              data.BodyCabineDamage,
		"autoCleanliness":               data.AutoCleanliness,
		"overallInteriorCleanliness":    data.OverallInteriorCleanliness,
		"stateCeling":                   data.StateCeling,
		"statePassengerSeat":            data.StatePassengerSeat,
		"stateDriverSeat":               data.StateDriverSeat,
		"stateSeatbelt":                 data.StateSeatbelt,
		"stateSteeringWheelAndSwitches": data.StateSteeringWheelAndSwitches,
		"statePanel":                    data.StatePanel,
		"stateSwitchKPP":                data.StateSwitchKPP,
		"windscreenCondition":           data.WindscreenCondition,
		"stateLeftwindscreen":           data.StateLeftwindscreen,
		"trunkCondition":                data.TrunkCondition,
		"stateTyre":                     data.StateTyre,
		"foreginLicenceRegistration":    data.ForeginLicenceRegistration,
		"dateData":                      data.DateData,
		"parts":                         bMPartsStruct,
		"otherParts":                    bMstrOtherPartsStruct,
		"token":                         data.Token,
		"free":                          data.Free,
		"dateInsert":                    time.Now(),
		"dateUpdate":                    nil,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateTechnicalServiceData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data DataTechnicalService
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	strPartsStruct, _ := data.Parts.(string)
	strOtherPartsStruct, _ := data.OtherParts.(string)
	strCoastJobsStruct, _ := data.CoastJobs.(string)

	var i []PartsStruct
	var y []OtherPartsStruct
	var z []CoastJobsStruct

	var bMPartsStruct []bson.M
	var bMstrOtherPartsStruct []bson.M
	var bMstrCoastJobsStruct []bson.M

	if err := json.Unmarshal([]byte(strPartsStruct), &i); err != nil {
		fmt.Println("ugh: ", err)
	}

	if err := json.Unmarshal([]byte(strOtherPartsStruct), &y); err != nil {
		fmt.Println("ugh: ", err)
	}

	if err := json.Unmarshal([]byte(strCoastJobsStruct), &z); err != nil {
		fmt.Println("ugh: ", err)
	}

	for _, u := range i {
		bMPartsStruct = append(bMPartsStruct, bson.M{
			"id":          u.Id,
			"title":       u.Title,
			"description": u.Description,
			"percent":     u.Percent,
			"summEnd":     u.SummEnd,
			"summ":        u.Summ,
		})
	}
	for _, u := range y {
		bMstrOtherPartsStruct = append(bMstrOtherPartsStruct, bson.M{
			"name":     u.Name,
			"articale": u.Articale,
			"howMuch":  u.HowMuch,
			"price":    u.Price,
			"percent":  u.Percent,
			"summEnd":  u.SummEnd,
		})
	}
	for _, u := range z {
		bMstrCoastJobsStruct = append(bMstrCoastJobsStruct, bson.M{
			"title":   u.Title,
			"price":   u.Price,
			"percent": u.Percent,
		})
	}
	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}
	podcastsCollection := mc.db.Collection("technicalService")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resultUpdate, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"auto":                          data.Auto,
				"typeJob":                       data.TypeJob,
				"autoRun":                       data.AutoRun,
				"coastJobsAll":                  data.CoastJobsAll,
				"contragent":                    data.Contragent,
				"statestatePassengerSeat":       data.StatestatePassengerSeat,
				"resultDyagnostic":              data.ResultDyagnostic,
				"coastSparePart":                data.CoastSparePart,
				"coastJobs":                     bMstrCoastJobsStruct,
				"statusRes":                     data.StatusRes,
				"tyreBrand":                     data.TyreBrand,
				"bodyCabineDamage":              data.BodyCabineDamage,
				"autoCleanliness":               data.AutoCleanliness,
				"overallInteriorCleanliness":    data.OverallInteriorCleanliness,
				"stateCeling":                   data.StateCeling,
				"statePassengerSeat":            data.StatePassengerSeat,
				"stateDriverSeat":               data.StateDriverSeat,
				"stateSeatbelt":                 data.StateSeatbelt,
				"stateSteeringWheelAndSwitches": data.StateSteeringWheelAndSwitches,
				"statePanel":                    data.StatePanel,
				"stateSwitchKPP":                data.StateSwitchKPP,
				"windscreenCondition":           data.WindscreenCondition,
				"stateLeftwindscreen":           data.StateLeftwindscreen,
				"trunkCondition":                data.TrunkCondition,
				"stateTyre":                     data.StateTyre,
				"foreginLicenceRegistration":    data.ForeginLicenceRegistration,
				"dateData":                      data.DateData,
				"parts":                         bMPartsStruct,
				"otherParts":                    bMstrOtherPartsStruct,
				"token":                         data.Token,
				"free":                          data.Free,
				"dateUpdate":                    time.Now(),
			},
		},
	)
	fmt.Println(resultUpdate.ModifiedCount) // output: 1
}

func (mc *MyClient) selectTechnicalServiceData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("technicalService")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.ParseForm()
	token := string(r.Form.Get("token"))
	cur, err := podcastsCollection.Find(ctx, bson.D{{"token", token}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData []DataTechnicalService
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		idJson, err := json.Marshal(result["_id"])
		autoJson, err := json.Marshal(result["auto"])
		typeJobJson, err := json.Marshal(result["typeJob"])
		autoRunJson, err := json.Marshal(result["autoRun"])
		coastJobsAllJson, err := json.Marshal(result["coastJobsAll"])
		contragentJson, err := json.Marshal(result["contragent"])
		statestatePassengerSeatJson, err := json.Marshal(result["statestatePassengerSeat"])
		resultDyagnosticJson, err := json.Marshal(result["resultDyagnostic"])
		coastSparePartJson, err := json.Marshal(result["coastSparePart"])
		coastJobsJson, err := json.Marshal(result["coastJobs"])
		statusResJson, err := json.Marshal(result["statusRes"])
		tyreBrandJson, err := json.Marshal(result["tyreBrand"])
		bodyCabineDamageJson, err := json.Marshal(result["bodyCabineDamage"])
		autoCleanlinessJson, err := json.Marshal(result["autoCleanliness"])
		overallInteriorCleanlinessJson, err := json.Marshal(result["overallInteriorCleanliness"])
		stateCelingJson, err := json.Marshal(result["stateCeling"])
		statePassengerSeatJson, err := json.Marshal(result["statePassengerSeat"])
		stateDriverSeatJson, err := json.Marshal(result["stateDriverSeat"])
		stateSeatbeltJson, err := json.Marshal(result["stateSeatbelt"])
		stateSteeringWheelAndSwitchesJson, err := json.Marshal(result["stateSteeringWheelAndSwitches"])
		statePanelJson, err := json.Marshal(result["statePanel"])
		stateSwitchKPPJson, err := json.Marshal(result["stateSwitchKPP"])
		windscreenConditionJson, err := json.Marshal(result["windscreenCondition"])
		stateLeftwindscreenJson, err := json.Marshal(result["stateLeftwindscreen"])
		trunkConditionJson, err := json.Marshal(result["trunkCondition"])
		stateTyreJson, err := json.Marshal(result["stateTyre"])
		dateDataJson, err := json.Marshal(result["dateData"])
		foreginLicenceRegistrationJson, err := json.Marshal(result["stateTyre"])

		idStr, _ := strconv.Unquote(string(idJson))
		autoStr, _ := strconv.Unquote(string(autoJson))
		typeJobStr, _ := strconv.Unquote(string(typeJobJson))
		autoRunStr, _ := strconv.Unquote(string(autoRunJson))
		coastJobsAllStr, _ := strconv.Unquote(string(coastJobsAllJson))
		contragentStr, _ := strconv.Unquote(string(contragentJson))
		statestatePassengerSeatStr, _ := strconv.Unquote(string(statestatePassengerSeatJson))
		resultDyagnosticStr, _ := strconv.Unquote(string(resultDyagnosticJson))
		coastSparePartStr, _ := strconv.Unquote(string(coastSparePartJson))
		coastJobsStr, _ := strconv.Unquote(string(coastJobsJson))
		statusResStr, _ := strconv.Unquote(string(statusResJson))
		tyreBrandStr, _ := strconv.Unquote(string(tyreBrandJson))
		bodyCabineDamageStr, _ := strconv.Unquote(string(bodyCabineDamageJson))
		autoCleanlinessStr, _ := strconv.Unquote(string(autoCleanlinessJson))
		overallInteriorCleanlinessStr, _ := strconv.Unquote(string(overallInteriorCleanlinessJson))
		stateCelingStr, _ := strconv.Unquote(string(stateCelingJson))
		statePassengerSeatStr, _ := strconv.Unquote(string(statePassengerSeatJson))
		stateDriverSeatStr, _ := strconv.Unquote(string(stateDriverSeatJson))
		stateSeatbeltStr, _ := strconv.Unquote(string(stateSeatbeltJson))
		stateSteeringWheelAndSwitchesStr, _ := strconv.Unquote(string(stateSteeringWheelAndSwitchesJson))
		statePanelStr, _ := strconv.Unquote(string(statePanelJson))
		stateSwitchKPPStr, _ := strconv.Unquote(string(stateSwitchKPPJson))
		windscreenConditionStr, _ := strconv.Unquote(string(windscreenConditionJson))
		stateLeftwindscreenStr, _ := strconv.Unquote(string(stateLeftwindscreenJson))
		trunkConditionStr, _ := strconv.Unquote(string(trunkConditionJson))
		stateTyreStr, _ := strconv.Unquote(string(stateTyreJson))
		dateDataStr, _ := strconv.Unquote(string(dateDataJson))
		foreginLicenceRegistrationStr, _ := strconv.ParseBool(string(foreginLicenceRegistrationJson))

		parsedData = append(parsedData, DataTechnicalService{
			Id:                            string(idStr),
			Auto:                          string(autoStr),
			TypeJob:                       string(typeJobStr),
			AutoRun:                       string(autoRunStr),
			CoastJobsAll:                  string(coastJobsAllStr),
			Contragent:                    string(contragentStr),
			StatestatePassengerSeat:       string(statestatePassengerSeatStr),
			ResultDyagnostic:              string(resultDyagnosticStr),
			CoastSparePart:                string(coastSparePartStr),
			CoastJobs:                     string(coastJobsStr),
			StatusRes:                     string(statusResStr),
			TyreBrand:                     string(tyreBrandStr),
			BodyCabineDamage:              string(bodyCabineDamageStr),
			AutoCleanliness:               string(autoCleanlinessStr),
			OverallInteriorCleanliness:    string(overallInteriorCleanlinessStr),
			StateCeling:                   string(stateCelingStr),
			StatePassengerSeat:            string(statePassengerSeatStr),
			StateDriverSeat:               string(stateDriverSeatStr),
			StateSeatbelt:                 string(stateSeatbeltStr),
			StateSteeringWheelAndSwitches: string(stateSteeringWheelAndSwitchesStr),
			StatePanel:                    string(statePanelStr),
			StateSwitchKPP:                string(stateSwitchKPPStr),
			WindscreenCondition:           string(windscreenConditionStr),
			StateLeftwindscreen:           string(stateLeftwindscreenStr),
			TrunkCondition:                string(trunkConditionStr),
			StateTyre:                     string(stateTyreStr),
			ForeginLicenceRegistration:    bool(foreginLicenceRegistrationStr),
			DateData:                      string(dateDataStr),
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))
}

func (mc *MyClient) deleteTechnicalServiceData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data DataTechnicalService
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}
	podcastsCollection := mc.db.Collection("technicalService")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}

func (mc *MyClient) selectTechnicalServiceDataOne(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var data DataTechnicalServicex
	r.ParseForm()
	idGet := string(r.Form.Get("id"))
	id, err := primitive.ObjectIDFromHex(strings.Trim(idGet, "\""))
	if err != nil {
		fmt.Println(err)
	}

	podcastsCollection := mc.db.Collection("technicalService")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.FindOne(
		ctx,
		bson.M{"_id": id}).Decode(&data)

	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)

	w.Write([]byte(bytes))
}
