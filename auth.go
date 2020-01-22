package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type AuthData struct {
	Id       string `json:"_id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
	Browser  string `json:"browser"`
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func randStringBytes() string {
	rand.Seed(time.Now().UTC().UnixNano())
	var n int
	n = rand.Intn(100)
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (mc *MyClient) insertAuthData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data AuthData
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

	podcastsCollection := mc.db.Collection("auth")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = podcastsCollection.InsertOne(ctx, bson.D{
		{"email", data.Email},
		{"phone", data.Phone},
		{"password", string(hashedPassword)},
		{"ip", data.Ip},
		{"browser", data.Browser},
		{"token", "fw54654wef6wef5w4EFFef_eff545"},
		{"dateInsert", time.Now()},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MyClient) updateAuthData(w http.ResponseWriter, r *http.Request) {
	fmt.Println(4)
}

func (mc *MyClient) checkAuthData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data AuthData
	err = json.NewDecoder(r.Body).Decode(&data)
	// Comparing the password with the hash

	podcastsCollection := mc.db.Collection("auth")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := podcastsCollection.Find(ctx, bson.M{
		"phone": data.Phone,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData AuthData
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		id, err := json.Marshal(result["_id"])
		password, err := json.Marshal(result["password"])

		sid, _ := strconv.Unquote(string(id))
		spassword, _ := strconv.Unquote(string(password))
		parsedData = AuthData{
			Id:       string(sid),
			Password: string(spassword),
		}

	}
	w.Header().Set("Content-Type", "application/json")

	err = bcrypt.CompareHashAndPassword([]byte(parsedData.Password), []byte(data.Password))
	if err != nil {
		w.Write([]byte("0"))
		return
	}
	bytes, err := json.Marshal(parsedData.Id)

	w.Write([]byte(bytes))

}

func (mc *MyClient) resetAuthData(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	var err error
	var data AuthData
	err = json.NewDecoder(r.Body).Decode(&data)
	// Comparing the password with the hash

	podcastsCollection := mc.db.Collection("auth")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := podcastsCollection.Find(ctx, bson.M{
		"email": data.Email,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var parsedData AuthData
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		id, err := json.Marshal(result["_id"])
		phone, err := json.Marshal(result["phone"])

		sid, _ := strconv.Unquote(string(id))
		sphone, _ := strconv.Unquote(string(phone))
		parsedData = AuthData{
			Id:    string(sid),
			Phone: string(sphone),
		}

	}
	if parsedData.Id == "" {
		w.Write([]byte("0"))
		return
	}
	var newPassword string
	newPassword = randStringBytes()

	password := []byte(newPassword)
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	resultUpdate, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"email": data.Email},
		bson.M{
			"$set": bson.M{
				"password":   string(hashedPassword),
				"dateUpdate": time.Now(),
			},
		},
	)
	fmt.Println(resultUpdate.ModifiedCount) // output: 1

	body := "Ваши новые доступы \n Логин: " + parsedData.Phone + "\n Пароль:" + newPassword
	header := make(map[string]string)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))
	auth := smtp.PlainAuth("", "comov.fedya@yandex.ru", "1xZZf405X58", "smtp.yandex.ru")
	err = smtp.SendMail(
		"smtp.yandex.ru:25",
		auth,
		"comov.fedya@yandex.ru",
		[]string{data.Email},
		[]byte(message))

	if err != nil {
		log.Fatal(err)
	}

}

func (mc *MyClient) deleteAuthData(w http.ResponseWriter, r *http.Request) {
	fmt.Println(4)
}
