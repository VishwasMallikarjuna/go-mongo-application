package mongoApi

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMongoCollection(collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}

func DatabaseHealthCheck(client *mongo.Collection) (string, string) {
	command := bson.D{{"dbStats", 1}}
	var result bson.D
	client.Database().RunCommand(context.TODO(), command).Decode(&result)
	fmt.Println(result)
	return fmt.Sprint(result[6].Value), fmt.Sprint(result[4].Value)
}
