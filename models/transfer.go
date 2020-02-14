package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Transfer struct {
	ID                     bson.ObjectId `bson:"_id" json:"id"`
	Account_Origin_Id      int           `bson:"account_origin_id" json:"account_origin_id"`
	Account_Destination_Id int           `bson:"account_destination_id" json:"account_destination_id"`
	Amount                 float32       `bson:"amount" json:"amount"`
	Created_At             time.Time     `bson:"created_at" json:"created_at"`
}
