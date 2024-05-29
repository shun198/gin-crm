package serializers

type LoginSerializer struct {
	EmployeeNumber string `json:"email" bson:"string"`
	Password       string `json:"password" bson:"password"`
}
