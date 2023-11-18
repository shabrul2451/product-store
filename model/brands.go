package model

import "time"

type Brand struct {
	Id        string    `bson:"id" json:"id"`
	Name      string    `json:"name" bson:"name"`
	StatusId  int       `bson:"status_id" json:"status_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
