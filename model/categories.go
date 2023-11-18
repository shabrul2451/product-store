package model

import "time"

type Category struct {
	Id        string    `bson:"id" json:"id"`
	Name      string    `json:"name" bson:"name"`
	ParentId  string    `json:"parent_id" bson:"parent_id"`
	Sequence  int64     `json:"sequence" bson:"sequence"`
	StatusId  int       `bson:"status_id" json:"status_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
