package api

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

type VacanciesResponse struct {
	Items   []Vacancy `json:"items"`
	Found   int       `json:"found"`
	Pages   int       `json:"pages"`
	Page    int       `json:"page"`
	PerPage int       `json:"per_page"`
}
