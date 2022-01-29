package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Recruitment struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Title       string        `bson:"title,omitempty"`
	Company     string        `bson:"company,omitempty"`
	Location    string        `bson:"location,omitempty"`
	Salary      string        `bson:"salary,omitempty"`
	UrlJob      string        `bson:"url_job,omitempty"`
	UrlCompany  string        `bson:"url_company,omitempty"`
	JobDeadline time.Time     `bson:"job_deadline,omitempty"`
}
