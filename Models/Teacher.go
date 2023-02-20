package Models

import (
	"fmt"
	"errors"
	"api/Config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllTeachers(teacher *[]Teacher) (err error) {
	if err = Config.DB.Find(teacher).Error; err != nil {
		return err;
	}
	return nil
}

func CreateTeacher(teacher *Teacher) (err error) {
	if err = Config.DB.Create(teacher).Error; err != nil {
		return errors.New("the specified teacher already exists");
	}
	return nil
}

func DeleteTeacher(teacher *Teacher) (err error) {
	if err = Config.DB.Model(&teacher).Where("email = ?", teacher.Email).First(&teacher).Error; err != nil {
		return fmt.Errorf("the specified teacher: %s does not exist in the database", teacher.Email);
	}

	if err = Config.DB.Where("email = ?", teacher.Email).Delete(teacher).Error; err != nil {
		return errors.New("an error occurred while attempting to delete the teacher")
	}
	return nil
}