package repository

import (
	"context"
	"log"
	service_skilltree "vacanciesParser/internal/core/logic/skilltree"
	"vacanciesParser/internal/core/repository/mongodb"
	mongodb_skilltree "vacanciesParser/internal/core/repository/mongodb/skilltree"
	mongodb_vacancy "vacanciesParser/internal/core/repository/mongodb/vacancy"
	"vacanciesParser/internal/core/repository/readonly/hh"
	"vacanciesParser/internal/core/repository/redis"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	client        *mongo.Client
	vacancyRepo   *mongodb_vacancy.Repository
	skilltreeRepo *mongodb_skilltree.Repository
}

// TODO: Текущая реализация пакета skilltree зависит от наличия в базе элемента под названием root.
// Либо убрать зависимость, либо создавать его автоматически при отсутствии.
func NewRepository() *Repository {
	client, err := mongodb.ConnectToMongoDB()
	if err != nil {
		log.Fatalf("ошибка подключения к Mongo DB: %v", err)
	}

	return &Repository{
		client:        client,
		vacancyRepo:   mongodb_vacancy.NewVacancyRepository(client),
		skilltreeRepo: mongodb_skilltree.NewSkillTreeRepository(client),
	}
}

func (r *Repository) Shutdown() {
	if err := r.client.Disconnect(context.Background()); err != nil {
		log.Printf("ошибка отключения от Mongo DB: %v", err)
	}
}

func (r *Repository) GetVacancies() {
	from := redis.GetLastExecutionDate()

	// Лучше сразу обновлять дату, чтобы не потерять вакансии, которые появятся в процессе работы программы.
	redis.UpdateLastExecutionDate()

	vacancies := make([]mongodb_vacancy.Vacancy, 0)
	vacancies = append(vacancies, hh.GetVacancies(from).ToMongo()...)

	if len(vacancies) == 0 {
		log.Printf("Новые вакансии на hh.ru не найдены.")
		return
	}

	docs := make([]any, len(vacancies))
	for i, data := range vacancies {
		docs[i] = data
	}

	r.vacancyRepo.InsertVacancies(context.TODO(), docs)
}

func (r *Repository) GetSkillTree(ctx context.Context) (*service_skilltree.Node, error) {
	return r.skilltreeRepo.GetSkillTree(ctx)
}

func (r *Repository) CreateNode(ctx context.Context, tree *service_skilltree.NodePath) error {
	return r.skilltreeRepo.CreateNode(ctx, tree)
}
