package Models

type Teacher struct {
	Email string `json:"email"`
}

type CreateTeacherInput struct {
	Email string `json:"email" binding:"required"`
}

type DeleteTeacherInput struct {
	Email string `json:"email" binding:"required"`
}

func (Teacher) TableName() string {
	return "Teachers"
}