package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type DataAllUsers struct {
	Id         string `json:"_id"`
	ParentId   string `json:"parentId"`
	Lastname   string `json:"lastname"`
	Firstname  string `json:"firstname"`
	Fathername string `json:"fathername"`
	TypePerson string `json:"type"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func (mc *MyClient) selectAllUsersData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var parsedData []DataAllUsers
	podcastsCollection := mc.db.Collection("allUsers")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.ParseForm()
	token := string(r.Form.Get("token"))
	cur, err := podcastsCollection.Find(ctx, bson.D{{"token", token}})
	//cur, err := podcastsCollection.Find(ctx, bson.D{})
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
		parentId, err := json.Marshal(result["parentId"])
		lastname, err := json.Marshal(result["lastname"])
		firstname, err := json.Marshal(result["firstname"])
		fathername, err := json.Marshal(result["fathername"])
		typePerson, err := json.Marshal(result["type"])
		email, err := json.Marshal(result["email"])

		sid, _ := strconv.Unquote(string(id))
		parentIdStr, _ := strconv.Unquote(string(parentId))
		lastnameStr, _ := strconv.Unquote(string(lastname))
		firstnameStr, _ := strconv.Unquote(string(firstname))
		fathernameStr, _ := strconv.Unquote(string(fathername))
		typePersonStr, _ := strconv.Unquote(string(typePerson))
		emailStr, _ := strconv.Unquote(string(email))

		parsedData = append(parsedData, DataAllUsers{
			Id:         string(sid),
			ParentId:   string(parentIdStr),
			Lastname:   string(lastnameStr),
			Firstname:  string(firstnameStr),
			Fathername: string(fathernameStr),
			TypePerson: string(typePersonStr),
			Email:      string(emailStr),
		})

	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(parsedData)

	w.Write([]byte(bytes))

}
