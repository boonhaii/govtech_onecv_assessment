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
		c.JSON(http.StatusNotFound, gin.H { "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H {
			"data": student,
		})
	}
}

func CreateStudent(c *gin.Context) {
	var input Models.CreateStudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"error": err.Error()})
	}
	
	student := Models.Student{Email: input.Email, Suspended: false}
	
	err := Models.CreateStudent(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H {
			"data": student,
			"message": fmt.Sprintf("Creation of %s was successful", student.Email),
		})
	}
}

func DeleteStudent(c *gin.Context) {
	var input Models.DeleteStudentInput
	if  err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"message": err.Error()})
	}

	student := Models.Student{Email: input.Email}
	err := Models.DeleteStudent(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H { 
			"message": fmt.Sprintf("Deletion of %s was successful", input.Email),
		})
	}
}

func SuspendStudent(c *gin.Context) {
	var input Models.SuspendStudentInput
	if  err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"message": err.Error()})
	}

	student := Models.Student{Email: input.Email}
	err := Models.SuspendStudent(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": err.Error(),
		})
	} else {
		c.Status(http.StatusNoContent)
	}
}