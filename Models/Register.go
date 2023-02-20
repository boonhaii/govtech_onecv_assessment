package Models

import (
	"fmt"
	"api/Config"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterStudent(teacher *Teacher, student *Student) (err error) {
	fmt.Println(teacher);
	fmt.Println(student)

	// if err = Config.DB.Model(&student).Find(student).Error; err != nil {
	// 	return err;
	// }

	// if err = Config.DB.Model(&teacher).Find(teacher).Error; err != nil {
	// 	return err;
	// }

	if err = Config.DB.Create(Register{T_email: teacher.Email, S_email: student.Email}).Error; err != nil {
		return err;
	}

	return nil
}