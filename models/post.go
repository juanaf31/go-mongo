package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetPosts struct{
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title,omitempty" bson:"title,omitempty"`
	Desc string `json:"desc,omitempty" bson:"desc,omitempty"`
	// Date primitive.DateTime `json:"date,omitempty" bson:"date,omitempty"`
}