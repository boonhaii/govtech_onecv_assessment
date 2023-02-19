package Controllers

import (
	"fmt"
	"net/http"
	"api/Models"
	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var student []Models.Student
	err := Models.GetAllStudents(&student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H {
			"data": student,
		})
	}
}

func CreateStudent(c *gin.Context) {
	var student Models.Student
	c.BindJSON(&student)
	fmt.Println(student)
	err := Models.CreateStudent(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Another student with the same email already exists!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H {
			"data": student,
			"message": fmt.Sprintf("Creation of %s was successful", student.Email),
		})
	}
}

func DeleteStudent(c *gin.Context) {
	var student Models.Student
	email := c.Params.ByName("email")
	err := Models.DeleteStudent(&student, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H { 
			"message": fmt.Sprintf("Deletion of %s was successful", email),
		})
	}
}

func SuspendStudent(c *gin.Context) {
	var student Models.Student
	c.BindJSON(&student)
	err := Models.SuspendStudent(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Student does not exist!",
		})
	} else {
		c.Status(http.StatusNoContent)
	}
}