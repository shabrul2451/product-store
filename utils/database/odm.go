package database

import (
	"Product-Store/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

type DmManager struct {
	Ctx context.Context
	Db  *mongo.Database
}

var singletonDmManager *DmManager
var onceDmManager sync.Once

// GetDmManager returns dmManager
func GetDmManager() *DmManager {
	onceDmManager.Do(func() {
		log.Println("[INFO] Starting Initializing Singleton DB Manager")
		singletonDmManager = &DmManager{}
		singletonDmManager.initConnection()
	})
	return singletonDmManager
}

func (dm *DmManager) initConnection() {
	ctx := context.Background()
	dm.Ctx = ctx
	clientOpts := options.Client().ApplyURI(config.DatabaseConnectionString)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Println("[ERROR] DB Connection error:", err.Error())
		return
	}
	db := client.Database(config.DatabaseName)
	dm.Db = db
	log.Println("[INFO] Initialized Singleton DB Manager")
}
