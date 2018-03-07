package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	."github.com/dongik/restapi/dao"
	."github.com/dongik/restapi/models"
	."github.com/dongik/restapi/config"
)

var config = Config{}
var dao = CardsDAO{}

func AllCardsEndpoint(w http.ResponseWriter, r *http.Request) {

	cards, err := dao.FindAll()
	if err != nil {
		fmt.Print("error")
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, cards)
}

func FindCardEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}


func UpdateCardEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteCardEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func CreateCardEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var card Card
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	card.ID = bson.NewObjectId()
	if err := dao.Insert(card); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, card)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/cards", AllCardsEndpoint).Methods("GET")
	r.HandleFunc("/cards", CreateCardEndpoint).Methods("POST")
	r.HandleFunc("/cards", UpdateCardEndpoint).Methods("PUT")
	r.HandleFunc("/cards", DeleteCardEndpoint).Methods("DELETE")
	r.HandleFunc("/cards/{id}", FindCardEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}