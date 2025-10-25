package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func GetITRolesIDs() []string {
	apiURL := "https://api.hh.ru/professional_roles"

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatalf("Ошибка при выполнении HTTP запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Неожиданный статус ответа: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении тела ответа: %v", err)
	}

	var data RolesResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Ошибка при десериализации JSON: %v", err)
	}

	for _, v := range data.Categories {
		if v.Name == "Информационные технологии" {
			return v.getRolesIDs()
		}
	}

	log.Fatalf("Ошибка: не найдена категория \"Информационные технологии\"")
	return nil
}

func BuildQuery(roles []string) string {
	baseURL := "https://api.hh.ru/vacancies"

	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}

	q := u.Query()

	q.Set("text", "(Go OR Golang) AND (NOT \"Яндекс GO\")") // Дурацкая доставка мешает нормально находить вакансии на Гошника 👺👺👺
	q.Set("search_field", "name")

	for _, v := range roles {
		q.Add("professional_role", v)
	}

	u.RawQuery = q.Encode()

	return u.String()
}

func GetVacancies(query string) VacanciesResponse {
	resp, err := http.Get(query)
	if err != nil {
		log.Fatalf("Ошибка при выполнении HTTP запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Неожиданный статус ответа: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении тела ответа: %v", err)
	}

	var data VacanciesResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Ошибка при десериализации JSON: %v", err)
	}

	return data
}
