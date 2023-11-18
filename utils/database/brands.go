package database

import (
	"Product-Store/interfaces/database"
	"Product-Store/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	BrandsCollection = "brands_collection"
)

type brandRepository struct {
	manager *DmManager
}

func (b brandRepository) GetById(id string) model.Brand {
	var res model.Brand
	query := bson.M{
		"$and": []bson.M{
			{"id": id},
		},
	}
	coll := b.manager.Db.Collection(BrandsCollection)
	result, err := coll.Find(b.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(model.Brand)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		res = *elemValue
	}
	return res
}

func (b brandRepository) Get() []model.Brand {
	var results []model.Brand
	query := bson.M{
		"$or": []interface{}{
			bson.M{"status_id": 1},
		},
	}
	coll := b.manager.Db.Collection(BrandsCollection)
	result, err := coll.Find(b.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(model.Brand)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

func (b brandRepository) Update(dto model.Brand) error {
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
	coll := b.manager.Db.Collection(BrandsCollection)
	err := coll.FindOneAndUpdate(b.manager.Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR] Insert document:", err.Err())
	}
	return nil
}

func (b brandRepository) Delete(id string) error {
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
	coll := b.manager.Db.Collection(BrandsCollection)
	err := coll.FindOneAndUpdate(b.manager.Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
}

func NewBrandRepository() database.BrandRepository {
	return &brandRepository{manager: GetDmManager()}
}
