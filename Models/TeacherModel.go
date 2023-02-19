package Models

type Teacher struct {
	Email string `json:"email"`
}

func (Teacher) TableName() string {
	return "Teachers"
}