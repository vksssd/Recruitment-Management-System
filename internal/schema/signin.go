package schema

// import (
// 	"time"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

type Signin struct {
		Email string `json:"email" bson:"email"`
		PasswordHash string `json:"password_hash" bson:"password_hash"`
}
