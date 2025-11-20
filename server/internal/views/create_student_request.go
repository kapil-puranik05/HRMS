package views

type CreateStudentRequest struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Department string `json:"department" binding:"required"`
}
