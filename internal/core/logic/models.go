package logic

import "time"

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
	Title           string
	Source          string
	URL             string
	Company         string
	MinPayment      int
	MaxPayment      int
	Currency        string
	WorkFormat      []string
	Experience      string
	PublicationDate time.Time
	IsProcessed     bool
}
