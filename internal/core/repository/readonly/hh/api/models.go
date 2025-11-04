package api

import (
	"log"
	"time"
	"vacanciesParser/internal/core/repository/mongodb"
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

type Salary struct {
	From     int    `json:"from"`
	To       int    `json:"to"`
	Currency string `json:"currency"`
	Gross    bool   `json:"gross"`
}

type Employer struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	CountryId    int    `json:"country_id"`
	IsAccredited bool   `json:"accredited_it_employer"`
}

type NameField struct {
	Name string `json:"name"`
}

type Vacancy struct {
	Name        string      `json:"name"`
	Salary      Salary      `json:"salary"`
	PublishedAt string      `json:"published_at"`
	URL         string      `json:"alternate_url"`
	Employer    Employer    `json:"employer"`
	WorkFormat  []NameField `json:"work_format"`
	Experience  NameField   `json:"experience"`
}

type Vacancies []Vacancy

type VacanciesResponse struct {
	Items   Vacancies `json:"items"`
	Found   int       `json:"found"`
	Pages   int       `json:"pages"`
	Page    int       `json:"page"`
	PerPage int       `json:"per_page"`
}

func (vs Vacancies) ToMongo() []mongodb.Vacancy {
	vacancies := make([]mongodb.Vacancy, 0, len(vs))

	for _, v := range vs {
		publicationDate, err := time.Parse(layout, v.PublishedAt)
		if err != nil {
			log.Printf("Error while parsing publication date: %v", err)
			publicationDate = time.UnixMicro(0)
		}

		workFormats := make([]string, 0)
		for _, format := range v.WorkFormat {
			workFormats = append(workFormats, format.Name)
		}

		var salary *mongodb.Salary
		if v.Salary.From == 0 && v.Salary.To == 0 {
			salary = nil
		} else {
			salary = &mongodb.Salary{
				From:     v.Salary.From,
				To:       v.Salary.To,
				Currency: v.Salary.Currency,
				Gross:    v.Salary.Gross,
			}
		}

		vacancies = append(vacancies, mongodb.Vacancy{
			Title:   v.Name,
			Source:  "hh",
			URL:     v.URL,
			Company: v.Employer.Name,
			Salary:  salary,
			Employer: &mongodb.Employer{
				Name:         v.Employer.Name,
				CountryId:    v.Employer.CountryId,
				IsAccredited: v.Employer.IsAccredited,
			},
			WorkFormat:      workFormats,
			Experience:      v.Experience.Name,
			PublicationDate: publicationDate,
			IsProcessed:     false,
		})
	}

	return vacancies
}
