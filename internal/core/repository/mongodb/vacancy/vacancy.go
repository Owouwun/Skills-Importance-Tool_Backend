package vacancy

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func getVacancyCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("skill_importance").Collection("vacancy")
}

func NewVacancyRepository(client *mongo.Client) *Repository {
	return &Repository{
		collection: getVacancyCollection(client),
	}
}

func (r *Repository) InsertVacancies(ctx context.Context, docs []interface{}) {
	result, err := r.collection.InsertMany(ctx, docs)
	if err != nil {
		log.Printf("ошибка при попытке вставки документов: %v", err)
		return
	}

	log.Printf("успешно вставлено документов: %d\n", len(result.InsertedIDs))
}
