package models

type UserType string

const (
	Applicant UserType = "Applicant"
	Admin     UserType = "Admin"
)