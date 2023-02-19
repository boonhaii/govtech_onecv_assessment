package Models

import (
	"api/Config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllStudents(student *[]Student) (err error) {
	if err = Config.DB.Find(student).Error; err != nil {
		return err;
	}
	return nil
}

func CreateStudent(student *Student) (err error) {
	if err = Config.DB.Create(student).Error; err != nil {
		return err;
	}
	return nil
}

// No SQL Injection risk since Golang precompiles the db query, and injection cannot occur.
func DeleteStudent(student *Student, email string) (err error) {
	if err = Config.DB.Where("email = ?", email).Delete(student).Error; err != nil {
		return err;
	}
	return nil
}

// Need to check if email actually exists.
func SuspendStudent(student *Student) (err error) {
	if err = Config.DB.Model(&student).Where("email = ?", student.Email).Update("suspended", true).Error; err != nil {
		return err;
	}
	return nil
}