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
	"golang.org/x/crypto/bcrypt"
)

type UserProfile struct {
	Id                 string `json:"_id"`
	Type               string `json:"type"`
	Email              string `json:"email"`
	Login              string `json:"login"`
	Password           string `json:"password"`
	LastName           string `json:"lastName"`
	Name               string `json:"name"`
	FatherName         string `json:"fatherName"`
	Phone              string `json:"phone"`
	StatusRes          string `json:"statusRes"`
	GroundsForContract string `json:"groundsForContract"`
	Token              string `json:"token"`
}

func (mc *MyClient) insertUserProfileData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data UserProfile
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	password := []byte(data.Password)
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	podcastsCollection := mc.db.Collection("userProfile")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"type", data.Type},
		{"email", data.Email},
		{"login", data.Login},
		{"password", string(hashedPassword)},
		{"lastName", data.LastName},
		{"name", data.Name},
		{"fatherName", data.FatherName},
		{"phone", data.Phone},
		{"token", data.Token},
		{"statusRes", data.StatusRes},
		{"groundsForContract", data.GroundsForContract},
		{"dateUpdate", time.Now()},
		{"dateInsert", time.Now()},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateUserProfileData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data UserProfile
	var bMAuth bson.M
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := primitive.ObjectIDFromHex(strings.Trim(data.Id, "\""))
	if err != nil {
		fmt.Println(err)
	}

	if data.Password != "" {
		password := []byte(data.Password)
		// Hashing the password with the default cost of 10
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		bMAuth = bson.M{
			"type":               data.Type,
			"email":              data.Email,
			"login":              data.Login,
			"password":           string(hashedPassword),
			"lastName":           data.LastName,
			"name":               data.Name,
			"fatherName":         data.FatherName,
			"phone":              data.Phone,
			"statusRes":          data.StatusRes,
			"groundsForContract": data.GroundsForContract,
			"dateUpdate":         time.Now(),
		}
	} else {
		bMAuth = bson.M{
			"type":               data.Type,
			"email":              data.Email,
			"login":              data.Login,
			"lastName":           data.LastName,
			"name":               data.Name,
			"fatherName":         data.FatherName,
			"phone":              data.Phone,
			"statusRes":          data.StatusRes,
			"groundsForContract": data.GroundsForContract,
			"dateUpdate":         time.Now(),
		}
	}

	podcastsCollection := mc.db.Collection("userProfile")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resultUpdate, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bMAuth,
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(resultUpdate.ModifiedCount) // output: 1
}

func (mc *MyClient) deleteUserProfileData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data UserProfile
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
	podcastsCollection := mc.db.Collection("userProfile")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	podcastsCollection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
}

func (mc *MyClient) selectUserProfileData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	podcastsCollection := mc.db.Collection("userProfile")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.ParseForm()
	token := string(r.Form.Get("token"))
	cur, err := podcastsCollection.Find(ctx, bson.D{{"token", token}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData []UserProfile
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		idJson, err := json.Marshal(result["_id"])
		typeJson, err := json.Marshal(result["type"])
		emailJson, err := json.Marshal(result["email"])
		loginJson, err := json.Marshal(result["login"])
		lastNameJson, err := json.Marshal(result["lastName"])
		nameJson, err := json.Marshal(result["name"])
		fatherNameJson, err := json.Marshal(result["fatherName"])
		phoneJson, err := json.Marshal(result["phone"])
		statusResJson, err := json.Marshal(result["statusRes"])
		groundsForContractJson, err := json.Marshal(result["groundsForContract"])

		idStr, _ := strconv.Unquote(string(idJson))
		typeStr, _ := strconv.Unquote(string(typeJson))
		emailStr, _ := strconv.Unquote(string(emailJson))
		loginStr, _ := strconv.Unquote(string(loginJson))
		lastNameStr, _ := strconv.Unquote(string(lastNameJson))
		nameStr, _ := strconv.Unquote(string(nameJson))
		fatherNameStr, _ := strconv.Unquote(string(fatherNameJson))
		phoneStr, _ := strconv.Unquote(string(phoneJson))
		statusResStr, _ := strconv.Unquote(string(statusResJson))
		groundsForContractStr, _ := strconv.Unquote(string(groundsForContractJson))

		parsedData = append(parsedData, UserProfile{
			Id:                 string(idStr),
			Type:               string(typeStr),
			Email:              string(emailStr),
			Login:              string(loginStr),
			LastName:           string(lastNameStr),
			Name:               string(nameStr),
			FatherName:         string(fatherNameStr),
			Phone:              string(phoneStr),
			StatusRes:          string(statusResStr),
			GroundsForContract: string(groundsForContractStr),
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))
}
