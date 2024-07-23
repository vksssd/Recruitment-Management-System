package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type CreateJob struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	JobId string `json:"job_id" bson:"job_id"`
	Title string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	PostedOn time.Time `json:"posted_on" bson:"posted_on"`
	TotalApplications int `json:"total_applications" bson:"total_applications"`
	CompanyName string `json:"company_name" bson:"company_name"`
	PostedBy User `json:"posted_by" bson:"posted_by"`	
}

type Job struct {
	Title string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	PostedOn time.Time `json:"posted_on" bson:"posted_on"`
	TotalApplications int `json:"total_applications" bson:"total_applications"`
	CompanyName string `json:"company_name" bson:"company_name"`
	PostedBy User `json:"posted_by" bson:"posted_by"`	
}