package dtos

type LoginDto struct {
	EmployeeNumber string `json:"email" bson:"string"`
	Password       string `json:"password" bson:"password"`
}
