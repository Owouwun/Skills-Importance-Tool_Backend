package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Salary struct {
	From     int
	To       int
	Currency string
	Gross    bool
}

type Employer struct {
	Name         string
	CountryId    int
	IsAccredited bool
}

type Vacancy struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Title           string             `bson:"title"`
	Source          string             `bson:"source"`
	URL             string             `bson:"url"`
	Company         string             `bson:"company"`
	Salary          *Salary            `bson:"salary,omitempty"`
	Employer        *Employer          `bson:"employer"`
	WorkFormat      []string           `bson:"work_format"`
	Experience      string             `bson:"experience"`
	PublicationDate time.Time          `bson:"publication_date"`
	IsProcessed     bool               `bson:"is_processed"`
}
