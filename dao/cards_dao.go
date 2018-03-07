package dao

import (
	"log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	."github.com/dongik/restapi/models"
)

type CardsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "cards"
)

func (m *CardsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}


func (m *CardsDAO) FindAll() ([]Card, error) {
	var cards []Card
	err := db.C(COLLECTION).Find(bson.M{}).All(&cards)
	return cards, err
}

func (m *CardsDAO) FindById(id string) (Card, error) {
	var card Card
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&card)
	return card, err
}

func (m *CardsDAO) Insert(card Card) error {
	err := db.C(COLLECTION).Insert(&card)
	return err
}

func (m *CardsDAO) Delete(card Card) error {
	err := db.C(COLLECTION).Remove(&card)
	return err
}

func (m *CardsDAO) Update(card Card) error {
	err := db.C(COLLECTION).UpdateId(card.ID, &card)
	return err
}

