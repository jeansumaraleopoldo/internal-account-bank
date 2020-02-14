package dao

import mgo "gopkg.in/mgo.v2"

var db *mgo.Database

const (
	COLLECTION_ACCOUNTS  = "accounts"
	COLLECTION_TRANSFERS = "transfers"
)
