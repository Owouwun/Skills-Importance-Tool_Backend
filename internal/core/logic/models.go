package logic

type VacancySalary struct {
	From     int
	To       int
	Currency string
	Gross    bool
}

type VacancyEmployer struct {
	ID           string
	Name         string
	CountryId    int
	IsAccredited bool
}

type Vacancy struct {
	Name        string
	Salary      VacancySalary
	PublishedAt string
	URL         string
	Employer    VacancyEmployer
}
