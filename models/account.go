package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"name" json:"name"`
	Cpf        string        `bson:"cpf" json:"cpf"`
	Balance    float32       `bson:"balance" json:"balance"`
	Created_At time.Time     `bson:"created_at" json:"created_at"`
}
