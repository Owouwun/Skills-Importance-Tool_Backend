package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB() (*mongo.Client, error) {
	return mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
}

func GetIDs(ctx context.Context, cursor *mongo.Cursor) map[primitive.ObjectID]struct{} {
	ids := make(map[primitive.ObjectID]struct{})

	for cursor.Next(ctx) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			log.Printf("ошибка при декодировании документа: %v", err)
			continue
		}

		ids[doc["_id"].(primitive.ObjectID)] = struct{}{}
	}

	return ids
}
