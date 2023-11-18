package database

import (
	"Product-Store/interfaces/database"
	"Product-Store/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type supplierRepository struct {
	manager *DmManager
}

func (s supplierRepository) GetById(id string) model.Supplier {
	var res model.Supplier
	query := bson.M{
		"$and": []bson.M{
			{"id": id},
		},
	}
	coll := s.manager.Db.Collection(suppliersCollection)
	result, err := coll.Find(s.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(model.Supplier)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		res = *elemValue
	}
	return res
}

var suppliersCollection = "suppliers_collection"

func (s supplierRepository) Create(dto model.Supplier) error {
	coll := s.manager.Db.Collection(suppliersCollection)
	_, err := coll.InsertOne(s.manager.Ctx, dto, nil)
	if err != nil {
		log.Println(err.Error())
	}
	return nil
}

func (s supplierRepository) Get() []model.Supplier {
	var results []model.Supplier
	query := bson.M{
		"$or": []interface{}{
			bson.M{"status_id": 1},
		},
	}
	coll := s.manager.Db.Collection(suppliersCollection)
	result, err := coll.Find(s.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(model.Supplier)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

func (s supplierRepository) Update(dto model.Supplier) error {
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
	coll := s.manager.Db.Collection(suppliersCollection)
	err := coll.FindOneAndUpdate(s.manager.Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR] Insert document:", err.Err())
	}
	return nil
}

func (s supplierRepository) Delete(id string) error {
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
	coll := s.manager.Db.Collection(suppliersCollection)
	err := coll.FindOneAndUpdate(s.manager.Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
}

func NewSupplierRepository() database.SupplierRepository {
	return &supplierRepository{manager: GetDmManager()}
}
