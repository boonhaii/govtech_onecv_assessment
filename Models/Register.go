package Models

import (
	"fmt"
	"errors"
	"api/Config"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterStudent(teacher *Teacher, student *Student) (err error) {
	if err = Config.DB.Model(&teacher).Where("email = ?", teacher.Email).First(&teacher).Error; err != nil {
		return errors.New("the specified teacher does not exist in the database");
	}

	if err = Config.DB.Model(&student).Where("email = ?", student.Email).First(&student).Error; err != nil {
		return fmt.Errorf("the specified student: %s does not exist in the database", student.Email);
	}

	if err = Config.DB.Create(Register{T_email: teacher.Email, S_email: student.Email}).Error; err != nil {
		return fmt.Errorf("an error occurred while registering the student: %s, did you already register the student?", student.Email);
	}

	return nil
}

func GetRegisteredStudents(teacher *Teacher, s_emails *[]string) (err error) {
	var registeredStudents []Register
	if err = Config.DB.Model(&teacher).Where("email = ?", teacher.Email).First(&teacher).Error; err != nil {
		return errors.New("the specified teacher does not exist in the database");
	}

	if err = Config.DB.Where("t_email = ?", teacher.Email).Find(&registeredStudents).Pluck("s_email", s_emails).Error; err != nil {
		return errors.New("an error occurred while finding the registered students");
	}
	
	return nil;
}
