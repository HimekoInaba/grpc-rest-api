package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password"`
	Email string `bson:"email" json:"email"`
	BirthDate string `bson:"birthDate" json:"birth_date"`
}

type LoginData struct {
	Email string `json:"email"`
	Password string `json:"password"`
}