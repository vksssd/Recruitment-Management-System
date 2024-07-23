package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type Applicant struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	ApplicantId string `json:"applicant_id" bson:"applicant_id"`
	Name	string `json:"name" bson:"name"`
	Email	string `json:"email" bson:"email"`
	Skills  string `json:"skills" bson:"skills"`
	ResumeFileAddr string `json:"resume_file_addr" bson:"resume_file_addr"`
	Education 	string `json:"education" bson:"education"`
	Experience  string `json:"experience" bson:"experience"`
	Phone       string `json:"phone" bson:"phone"`
}

type Profile struct {
	Name	string `json:"name" bson:"name"`
	Email	string `json:"email" bson:"email"`
	Skills  string `json:"skills" bson:"skills"`
	ResumeFileAddr string `json:"resume_file_addr" bson:"resume_file_addr"`
	Education 	string `json:"education" bson:"education"`
	Experience  string `json:"experience" bson:"experience"`
	Phone       string `json:"phone" bson:"phone"`
}

