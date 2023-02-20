package Models

type Student struct {
	Email string `json:"email"`
	Suspended bool `json:"suspended"`
}

type CreateStudentInput struct {
	Email string `json:"email" binding:"required"`
}

type DeleteStudentInput struct {
	Email string `json:"email" binding:"required"`
}

type SuspendStudentInput struct {
	Email string `json:"email" binding:"required"`
}

func (Student) TableName() string {
	return "Students"
}