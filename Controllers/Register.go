package Controllers

import (
	"net/http"
	"api/Models"
	"github.com/gin-gonic/gin"
)

func RegisterStudent(c *gin.Context) {
	var input Models.RegisterStudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"error": err.Error()})
	}

	teacher := Models.Teacher{Email: input.T_email}
	for _, email := range input.Students {
		student := Models.Student{Email: email}
		err := Models.RegisterStudent(&teacher, &student)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}
	}

	c.Status(http.StatusNoContent)
}