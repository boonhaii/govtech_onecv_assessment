package Controllers

import (
	"net/http"
	"api/Models"
	"api/Util"
	"github.com/gin-gonic/gin"
	"regexp"
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
		commonStudents = Util.Intersection(commonStudents, studentList)
	}

	c.JSON(http.StatusOK, gin.H {
		"students": commonStudents,
	})
}

// Need to check if the student actually exists.
func RetrieveStudentsForNotification(c *gin.Context) {
	var input Models.RetrieveStudentsForNotificationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"error": err.Error()})
	}

	teacher := Models.Teacher{ Email: input.T_email }
	var toNotify []string
	var suspended []string
	Models.GetRegisteredStudents(&teacher, &toNotify)
	Models.GetSuspendedStudents(&suspended)

	pattern := `\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b`
	re := regexp.MustCompile(pattern)
	tagged := re.FindAllString(input.Notification, -1)

	toNotify = append(toNotify, tagged...)
	toNotify = Util.Difference(toNotify, suspended)

	c.JSON(http.StatusOK, gin.H {
		"recipients": toNotify,
	})

}
