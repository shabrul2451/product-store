package database

import (
	"Product-Store/common"
	"Product-Store/interfaces/database"
	"Product-Store/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type productRepository struct {
	manager *DmManager
}

func (p productRepository) GetByFilter(filter common.ProductFilter) ([]model.Product, int64) {
	var queryVal bson.M
	var results []model.Product
	var query []bson.M
	query = append(query, bson.M{"status_id": 1})
	if filter.Name != "" {
		query = append(query, bson.M{"name": filter.Name})
	} else if filter.Category != "" {
		query = append(query, bson.M{"category_id": filter.Category})
	} else if filter.Supplier != "" {
		query = append(query, bson.M{"supplier_id": filter.Supplier})
	} else if filter.MaxPrice > 0 {
		query = append(query, bson.M{"unit_price": bson.M{"$lte": filter.MaxPrice}})
	} else if filter.MinPrice > 0 {
		query = append(query, bson.M{"unit_price": bson.M{"$gte": filter.MinPrice}})
	} else if filter.Brands != nil {
		query = append(query, bson.M{"brand_id": bson.M{"$in": filter.Brands}})
	}
	if len(query) > 0 {
		queryVal = bson.M{
			"$and": query,
		}
	}
	coll := p.manager.Db.Collection(productCollection)
	skip := filter.Page * filter.Limit
	result, err := coll.Find(p.manager.Ctx, queryVal, &options.FindOptions{
		Limit: &filter.Limit,
		Skip:  &skip,
		Sort:  bson.M{"unit_price": 1},
	})
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(model.Product)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	count, err := coll.CountDocuments(p.manager.Ctx, queryVal)
	if err != nil {
		log.Println(err)
	}
	return results, count
}

func (p productRepository) GetById(id string) model.Product {
	var res model.Product
	query := bson.M{
		"$and": []bson.M{
			{"supplier_id": id},
			{"status_id": 1},
		},
	}
	coll := p.manager.Db.Collection(productCollection)
	result, err := coll.Find(p.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(model.Product)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		res = *elemValue
	}
	return res
}

func (p productRepository) GetBySupplierIdAndProductName(id, name string) model.Product {
	var res model.Product
	query := bson.M{
		"$and": []bson.M{
			{"supplier_id": id},
			{"status_id": 1},
			{"name": name},
		},
	}
	coll := p.manager.Db.Collection(productCollection)
	result, err := coll.Find(p.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(model.Product)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		res = *elemValue
	}
	return res
}

var productCollection = "product_collection"

func (p productRepository) Create(dto model.Product) error {
	coll := p.manager.Db.Collection(productCollection)
	_, err := coll.InsertOne(p.manager.Ctx, dto, nil)
	if err != nil {
		log.Println(err.Error())
	}
	return nil
}

func (p productRepository) Get() []model.Product {
	var results []model.Product
	query := bson.M{
		"$or": []interface{}{
			bson.M{"status_id": 1},
		},
	}
	coll := p.manager.Db.Collection(productCollection)
	result, err := coll.Find(p.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(model.Product)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

func (p productRepository) Update(dto model.Product) error {
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
	coll := p.manager.Db.Collection(productCollection)
	err := coll.FindOneAndUpdate(p.manager.Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR] Insert document:", err.Err())
	}
	return nil
}

func (p productRepository) Delete(id string) error {
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
	coll := p.manager.Db.Collection(productCollection)
	err := coll.FindOneAndUpdate(p.manager.Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
}

func NewProductRepository() database.ProductRepository {
	return &productRepository{manager: GetDmManager()}
}
