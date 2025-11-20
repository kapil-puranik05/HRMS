package views

type UpdateStudentRequest struct {
	Name       *string `json:"name"`
	Email      *string `json:"email" binding:"omitempty,email"`
	Department *string `json:"department"`
}
