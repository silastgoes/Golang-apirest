package models

import "gopkg.in/mgo.v2/bson"

type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Genre       string        `bson:"genre" json:"genre"`
	Description string        `bson:"description" json:"description"`
}
