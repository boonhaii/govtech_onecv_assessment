package Models

type Register struct {
	T_email string `json:"t_email"`
	S_email string `json:"s_email"`
}

func (Register) TableName() string {
	return "Registers"
}