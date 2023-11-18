package model

import "time"

type Supplier struct {
	Id                 string    `bson:"id" json:"id"`
	Name               string    `bson:"name" json:"name"`
	Email              string    `bson:"email" json:"email"`
	Phone              string    `bson:"phone" json:"phone"`
	StatusId           int       `bson:"status_id" json:"status_id"`
	IsVerifiedSupplier bool      `bson:"is_verified_supplier" json:"is_verified_supplier"`
	CreatedAt          time.Time `bson:"created_at" json:"created_at"`
}
