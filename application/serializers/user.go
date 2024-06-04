package serializers

type ChangeUserDetailsSerializer struct {
	Name  string `json:"name" validate:"omitempty,max=255"`
	Email string `json:"email" validate:"omitempty,email,max=254"`
	Role  string `json:"role" validate:"omitempty,oneof=管理者 一般"`
}

type SendInviteUserEmailSerializer struct {
	Name           string `json:"name" validate:"required,max=255"`
	EmployeeNumber string `json:"employee_number" validate:"required,len=8,numeric"`
	Email          string `json:"email" validate:"required,email,max=254"`
	Role           string `json:"role" validate:"required,oneof=管理者 一般"`
}

type VerifyUserSerializer struct {
	Token           string `json:"token" validate:"required,max=255"`
	NewPassword     string `json:"new_password" validate:"required,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required,max=255"`
}

type SendResetPasswordEmailSerializer struct {
	Email string `json:"email" validate:"required,email,max=254"`
}

type ChangePasswordSerializer struct {
	CurrentPassword string `json:"current_password" validate:"required,max=255"`
	NewPassword     string `json:"new_password" validate:"required,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required,max=255"`
}

type ResetPasswordSerializer struct {
	Token           string `json:"token" validate:"required,max=255"`
	NewPassword     string `json:"new_password" validate:"required,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required,max=255"`
}

type CheckInvitationTokenSerializer struct {
	Token string `json:"token" validate:"required,max=255"`
}

type CheckResetPasswordTokenSerializer struct {
	Token string `json:"token" validate:"required,max=255"`
}
