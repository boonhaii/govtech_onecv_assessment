package Routes

import (
	"api/Controllers"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()
	teachers := router.Group("/api/teachers")
	{
		teachers.GET("", Controllers.GetTeachers)
		teachers.POST("", Controllers.CreateTeacher)
		teachers.DELETE("/:email", Controllers.DeleteTeacher)
	}

	students := router.Group("/api/students")
	{
		students.GET("", Controllers.GetStudents)
		students.POST("", Controllers.CreateStudent)
		students.DELETE("/:email", Controllers.DeleteStudent)
	}

	router.GET("/", func (c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to this API!")
	})

	// Route for teacher to register student (US.1)
	// router.POST("/api/register", Controllers.registerStudent)
	
	// // Route for retrieving common students between teachers (US.2)
	// router.GET("/api/commonstudents", Controllers.commonStudents)

	// // Route for teacher to suspend student (US.3)
	// // Should be under Student Model.
	router.POST("/api/suspend", Controllers.SuspendStudent)

	// // Route for list of student involved in notification (US. 4)
	// router.POST("/api/retrievefornotifications", Controllers.retrieveStudentsForNotification)

	return router
}