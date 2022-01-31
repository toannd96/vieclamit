package models

import (
	"time"
)

// Recruitment struct
type Recruitment struct {
	Title       string        `bson:"title,omitempty"`
	Company     string        `bson:"company,omitempty"`
	Location    string        `bson:"location,omitempty"`
	Salary      string        `bson:"salary,omitempty"`
	UrlJob      string        `bson:"url_job,omitempty"`
	UrlCompany  string        `bson:"url_company,omitempty"`
	JobDeadline time.Time     `bson:"job_deadline,omitempty"`
}

//Recruitments
type Recruitments []Recruitment
