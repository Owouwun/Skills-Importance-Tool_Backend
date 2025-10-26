package api

import (
	"log"
	"time"
	"vacanciesParser/internal/core/logic"
)

// API hh.ru говорит, что использует ISO 8601, но почему-то у них нет символа ":" в таймзоне
const layout = "2006-01-02T15:04:05Z0700"

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (cat Category) getRolesIDs() []string {
	ids := make([]string, 0)

	for _, v := range cat.Roles {
		ids = append(ids, v.ID)
	}

	return ids
}

type Category struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Roles []Role `json:"roles"`
}

type RolesResponse struct {
	Categories []Category `json:"categories"`
}

type VacancySalary struct {
	From     int    `json:"from"`
	To       int    `json:"to"`
	Currency string `json:"currency"`
	Gross    bool   `json:"gross"`
}

type VacancyEmployer struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	CountryId    int    `json:"country_id"`
	IsAccredited bool   `json:"accredited_it_employer"`
}

type Vacancy struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Salary      VacancySalary   `json:"salary"`
	PublishedAt string          `json:"published_at"`
	URL         string          `json:"alternate_url"`
	Employer    VacancyEmployer `json:"employer"`
}

type Vacancies []Vacancy

type VacanciesResponse struct {
	Items   Vacancies `json:"items"`
	Found   int       `json:"found"`
	Pages   int       `json:"pages"`
	Page    int       `json:"page"`
	PerPage int       `json:"per_page"`
}

func (vs Vacancies) ToLogic() []logic.Vacancy {
	vacancies := make([]logic.Vacancy, 0, len(vs))

	for _, v := range vs {
		publicationDate, err := time.Parse(layout, v.PublishedAt)
		if err != nil {
			log.Printf("Error while parsing publication date: %v", err)
			publicationDate = time.UnixMicro(0)
		}

		vacancies = append(vacancies, logic.Vacancy{
			Title:           v.Name,
			Source:          "hh",
			URL:             v.URL,
			Company:         v.Employer.Name,
			MinPayment:      v.Salary.From,
			MaxPayment:      v.Salary.To,
			Currency:        v.Salary.Currency,
			PublicationDate: publicationDate,
			IsProcessed:     false,
		})
	}

	return vacancies
}
