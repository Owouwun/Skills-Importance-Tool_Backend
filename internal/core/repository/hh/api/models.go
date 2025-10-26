package api

import (
	"vacanciesParser/internal/core/logic"
)

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
		vacancies = append(vacancies, logic.Vacancy{
			Name:        v.Name,
			Salary:      logic.VacancySalary(v.Salary),
			PublishedAt: v.PublishedAt,
			URL:         v.URL,
			Employer:    logic.VacancyEmployer(v.Employer),
		})
	}

	return vacancies
}
