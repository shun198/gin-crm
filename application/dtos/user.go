package dtos

type ChangeUserDetailsDto struct {
	Name  *string `json:"name" validate:"omitempty,max=255"`
	Email *string `json:"email" validate:"omitempty,email,max=254"`
	Role  *string `json:"role" validate:"omitempty,oneof=管理者 一般"`
}

type SendInviteUserEmailDto struct {
	Name           *string `json:"name" validate:"required,max=255"`
	EmployeeNumber *string `json:"employee_number" validate:"required,len=8,numeric"`
	Email          *string `json:"email" validate:"required,email,max=254"`
	Role           *string `json:"role" validate:"required,oneof=管理者 一般"`
}
