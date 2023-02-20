package Controllers

import (
	"fmt"
	"net/http"
	"api/Models"
	"github.com/gin-gonic/gin"
)

func GetTeachers(c *gin.Context) {
	var teacher []Models.Teacher
	err := Models.GetAllTeachers(&teacher)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H {
			"data": teacher,
		})
	}
}

func CreateTeacher(c *gin.Context) {
	var input Models.CreateTeacherInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"message": err.Error()})
	}
	
	teacher := Models.Teacher{Email: input.Email}
	
	err := Models.CreateTeacher(&teacher)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H {
			"data": teacher,
			"message": fmt.Sprintf("Creation of %s was successful", teacher.Email),
		})
	}
}

func DeleteTeacher(c *gin.Context) {
	var input Models.DeleteTeacherInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"message": err.Error()})
	}
	
	teacher := Models.Teacher{Email: input.Email}
	err := Models.DeleteTeacher(&teacher)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H { 
			"message": fmt.Sprintf("Deletion of %s was successful", input.Email),
		})
	}
}