package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"name" json:"name"`
	Cpf        int           `bson:"cpf" json:"cpf"`
	Ballance   float32       `bson:"ballance" json:"ballance"`
	Created_At time.Time     `bson:"created_at" json:"created_at"`
}
