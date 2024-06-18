package mongoApi

import "go.mongodb.org/mongo-driver/mongo"

func GetMongoCollection(collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}
