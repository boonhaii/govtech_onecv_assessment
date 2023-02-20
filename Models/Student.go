package Models

import (
	"fmt"
	"errors"
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
		return errors.New("the specified student already exists");
	}
	return nil
}

func DeleteStudent(student *Student) (err error) {
	if err = Config.DB.Model(&student).Where("email = ?", student.Email).First(&student).Error; err != nil {
		return fmt.Errorf("the specified student: %s does not exist in the database", student.Email)
	}

	if err = Config.DB.Model(&student).Where("email = ?", student.Email).Delete(student).Error; err != nil {
		return errors.New("an error occurred while attempting to delete the student")
	}
	return nil
}

// Need to check if email actually exists.
func SuspendStudent(student *Student) (err error) {
	if err = Config.DB.Model(&student).Where("email = ?", student.Email).First(&student).Error; err != nil {
		return fmt.Errorf("the specified student: %s does not exist in the database", student.Email);
	}

	if err = Config.DB.Model(&student).Where("email = ?", student.Email).Update("suspended", true).Error; err != nil {
		return err;
	}
	return nil
}

func GetSuspendedStudents(suspended *[]string) (err error) {
	var student Student
	if err = Config.DB.Model(&student).Where("suspended = true").Pluck("email", suspended).Error; err != nil {
		return err;
	}
	return nil
}

func FilterExistingStudents(s_emails *[]string) (err error) {
	var student Student
	if err = Config.DB.Model(&student).Where("email IN (?)", *s_emails).Pluck("email", s_emails).Error; err != nil {
		return err;
	}
	return nil
}