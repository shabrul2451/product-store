package model

import "time"

type ProductStock struct {
	Id            string    `bson:"id" json:"id"`
	ProductId     string    `bson:"product_id" json:"product_id"`
	StockQuantity int64     `bson:"stock_quantity" json:"stock_quantity"`
	UpdatedAt     time.Time `bson:"updated_at" json:"updated_at"`
}
