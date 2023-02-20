package Models

import (
	"api/Config"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterStudent(teacher *Teacher, student *Student) (err error) {
	if err = Config.DB.Model(&teacher).Where("email = ?", teacher.Email).Error; err != nil {
		return err;
	}

	if err = Config.DB.Model(&student).Where("email = ?", student.Email).Error; err != nil {
		return err;
	}

	if err = Config.DB.Create(Register{T_email: teacher.Email, S_email: student.Email}).Error; err != nil {
		return err;
	}

	return nil
}