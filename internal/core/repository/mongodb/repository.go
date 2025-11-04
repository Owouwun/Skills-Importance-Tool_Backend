package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InsertData(ctx context.Context, docs []interface{}) {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Не удалось подключиться к MongoDB: %v", err)
	}

	result, err := client.Database("vacancies_parser").Collection("vacancy").InsertMany(ctx, docs)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("Успешно вставлено документов: %d", len(result.InsertedIDs))
}
