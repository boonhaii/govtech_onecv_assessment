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
	var teacher Models.Teacher
	c.BindJSON(&teacher) // Deserializes binary into struct.
	err := Models.CreateTeacher(&teacher)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Another teacher with the same email already exists!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H {
			"data": teacher,
			"message": fmt.Sprintf("Creation of %s was successful", teacher.Email),
		})
	}
}

func DeleteTeacher(c *gin.Context) {
	var teacher Models.Teacher
	email := c.Params.ByName("email")
	err := Models.DeleteTeacher(&teacher, email)
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