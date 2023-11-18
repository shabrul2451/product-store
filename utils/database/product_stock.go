package database

import (
	"Product-Store/interfaces/database"
	"Product-Store/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type productStockRepository struct {
	manager *DmManager
}

func (p productStockRepository) GetById(id string) model.ProductStock {
	var res model.ProductStock
	query := bson.M{
		"$and": []bson.M{
			{"id": id},
		},
	}
	coll := p.manager.Db.Collection(productStockCollection)
	result, err := coll.Find(p.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(model.ProductStock)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		res = *elemValue
	}
	return res
}

var productStockCollection = "product_stock_collection"

func (p productStockRepository) Get() []model.ProductStock {
	var results []model.ProductStock
	query := bson.M{
		"$or": []interface{}{
			bson.M{"status_id": 1},
		},
	}
	coll := p.manager.Db.Collection(productStockCollection)
	result, err := coll.Find(p.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(model.ProductStock)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

func (p productStockRepository) Update(dto model.ProductStock) error {
	filter := bson.M{
		"$and": []bson.M{
			{"id": dto.Id},
		},
	}
	update := bson.M{
		"$set": dto,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := p.manager.Db.Collection(productStockCollection)
	err := coll.FindOneAndUpdate(p.manager.Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR] Insert document:", err.Err())
	}
	return nil
}

func (p productStockRepository) Delete(id string) error {
	filter := bson.M{
		"$and": []bson.M{
			{"id": id},
		},
	}
	update := bson.M{
		"$set": bson.M{"status_id": 0},
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := p.manager.Db.Collection(productStockCollection)
	err := coll.FindOneAndUpdate(p.manager.Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
}

func NewProductStockRepository() database.ProductStockRepository {
	return &productStockRepository{manager: GetDmManager()}
}
