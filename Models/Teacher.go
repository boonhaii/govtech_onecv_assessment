package Models

import (
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
		return err;
	}
	return nil
}

// No SQL Injection risk since Golang precompiles the db query, and injection cannot occur.
func DeleteTeacher(teacher *Teacher, email string) (err error) {
	if err = Config.DB.Where("email = ?", email).Delete(teacher).Error; err != nil {
		return err;
	}
	return nil
}