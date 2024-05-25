package dtos

type ChangeUserDetailsDto struct {
	Name  string `json:"name" validate:"required,min=2,max=255"`
	Email string `json:"email" validate:"required,email"`
	Role  string `json:"role" validate:"required,oneof=ADMIN GENERAL"`
}
