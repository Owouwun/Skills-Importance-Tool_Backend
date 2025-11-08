package main

import (
	"log"
	"vacanciesParser/internal/app"
	"vacanciesParser/internal/core/logic/vacancies"
	"vacanciesParser/internal/core/repository"
)

func main() {
	repo := repository.NewRepository()

	vs := vacancies.NewService(repo)
	vs.GetVacancies()

	router := app.PrepareRouter()

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
