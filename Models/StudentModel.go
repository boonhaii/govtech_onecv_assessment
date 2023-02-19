package Models

type Student struct {
	Email string `json:"email"`
	Suspended bool `json:"suspended"`
}

func (Student) TableName() string {
	return "Students"
}