package Models

type Teacher struct {
	Email string `json:"email"`
}

type CreateTeacherInput struct {
	Email string `json:"teacher" binding:"required"`
}

type DeleteTeacherInput struct {
	Email string `json:"teacher" binding:"required"`
}

func (Teacher) TableName() string {
	return "Teachers"
}