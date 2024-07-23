package models

type User struct {
	Name string
	Email string
	Address string
	UserType UserType //need work on refrencing
	PasswordHash string
	ProfileHeading string
	// Profle Profle  //need to work on refrencing
}