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
			return
		}
	}

	c.Status(http.StatusNoContent)
}

func CommonStudents(c *gin.Context) {
	t_emails, ok := c.GetQueryArray("teacher")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var registeredStudents [][]string
	for _, t_email := range t_emails {
		var students []string
		teacher := Models.Teacher{Email: t_email}
		Models.GetRegisteredStudents(&teacher, &students)
		registeredStudents = append(registeredStudents, students)
	}

	commonStudents := registeredStudents[0]
	for _, studentList := range registeredStudents {
		commonStudents = intersection(commonStudents, studentList)
	}

	c.JSON(http.StatusOK, gin.H {
		"students": commonStudents,
	})
}

// To move to util file
func intersection(s1, s2 []string) (res []string) {
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}

	for _, e := range s2 {
		if hash[e] {
			res = append(res, e)
		}
	}

	return res
}
