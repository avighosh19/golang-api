package models

import "gopkg.in/mgo.v2/bson"


type User struct{
	Id		bson.ObjectId   `json:"id" bson:"_id"`
	Name	string			`json:"name" bson:"name"`
	Gender	string			`json:"gender" bson:"gender"`
	Age      int			`json:"age" bson:"age"`
}


type Post struct {
	Id       bson.ObjectId       `json:"id" bson:"_id"`
	caption  string              `json:"caption" bson"caption"`
	imageurl string              `json:"imageurl" bson:"imageurl"`
	postTime bson.MongoTimestamp `json:"postTime" bson:"postTime"`
}
