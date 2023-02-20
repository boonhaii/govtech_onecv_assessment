package Models

type Register struct {
	T_email string `json:"t_email"`
	S_email string `json:"s_email"`
}

type RegisterStudentInput struct {
	T_email string `json:"teacher" binding:"required"`
	Students []string `json:"students" binding:"required"`
}

func (Register) TableName() string {
	return "Registers"
}