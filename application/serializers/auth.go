package serializers

type LoginSerializer struct {
	EmployeeNumber string `json:"employee_number" bson:"string"`
	Password       string `json:"password" bson:"password"`
}
