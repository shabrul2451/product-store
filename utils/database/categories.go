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
	CategoriesCollection = "categories_collection"
)

type categoryRepository struct {
	manager *DmManager
}

func (c categoryRepository) GetById(id string) model.Category {
	var res model.Category
	query := bson.M{
		"$and": []bson.M{
			{"id": id},
		},
	}
	coll := c.manager.Db.Collection(CategoriesCollection)
	result, err := coll.Find(c.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(model.Category)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		res = *elemValue
	}
	return res
}

func (c categoryRepository) Create(dto model.Category) error {
	cat := c.Get()
	if cat != nil {
		count := c.countDocuments()
		dto.Sequence = count + 1
	} else {
		dto.Sequence = 1
	}
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
	coll := c.manager.Db.Collection(CategoriesCollection)
	err := coll.FindOneAndUpdate(c.manager.Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR] Insert document:", err.Err())
	}
	return nil
}

func (c categoryRepository) Get() []model.Category {
	var results []model.Category
	query := bson.M{
		"$or": []interface{}{
			bson.M{"status_id": 1},
		},
	}
	coll := c.manager.Db.Collection(CategoriesCollection)
	result, err := coll.Find(c.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(model.Category)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

func (c categoryRepository) Update(dto model.Category) error {
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
	coll := c.manager.Db.Collection(CategoriesCollection)
	err := coll.FindOneAndUpdate(c.manager.Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR] Insert document:", err.Err())
	}
	return nil
}

func (c categoryRepository) Delete(id string) error {
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
	coll := c.manager.Db.Collection(CategoriesCollection)
	err := coll.FindOneAndUpdate(c.manager.Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
}

func (c categoryRepository) countDocuments() int64 {
	filter := bson.M{
		"$and": []bson.M{
			{"status_id": 1},
		},
	}
	coll := c.manager.Db.Collection(CategoriesCollection)
	count, err := coll.CountDocuments(c.manager.Ctx, filter)
	if err != nil {
		log.Println(err)
	}
	return count
}

func NewCategoriesRepository() database.CategoriesRepository {
	return &categoryRepository{manager: GetDmManager()}
}
