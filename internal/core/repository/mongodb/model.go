package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Salary struct {
	From     int    `bson:"from"`
	To       int    `bson:"to"`
	Currency string `bson:"currency"`
	IsGross  bool   `bson:"is_gross"`
}

type Employer struct {
	Name         string `bson:"name"`
	CountryId    int    `bson:"country_id"`
	IsAccredited bool   `bson:"is_accredited"`
}

// From и To равны 0, если не установлены или не релевантны
type ExperienceByYears struct {
	From int `bson:"from"`
	To   int `bson:"to"`
}

type Vacancy struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Title             string             `bson:"title"`
	Source            string             `bson:"source"`
	URL               string             `bson:"url"`
	Company           string             `bson:"company"`
	Salary            *Salary            `bson:"salary,omitempty"`
	Employer          *Employer          `bson:"employer"`
	WorkFormat        []string           `bson:"work_format"`
	ExperienceByYears ExperienceByYears  `bson:"experience_by_years"`
	PublicationDate   time.Time          `bson:"publication_date"`
	IsProcessed       bool               `bson:"is_processed"`
}

type Skill struct {
	ID       primitive.ObjectID  `bson:"_id"`
	Name     string              `bson:"name"`
	ParentID *primitive.ObjectID `bosn:"parent_id,omitempty"`
}
