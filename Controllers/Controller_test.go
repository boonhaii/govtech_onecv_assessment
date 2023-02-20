package Controllers

import (
	"api/Config"
	"api/Models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func setupDB() {
	var err error
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig("test"))) // To setup MySQL DB.

	if err != nil {
		fmt.Println("Status: ", err)
	}

	Config.DB.AutoMigrate(&Models.Teacher{}, &Models.Student{}, &Models.Register{})
	populateTestDB()
}

func clearDB() {
	Config.DB.Exec("DELETE FROM Teachers")
	Config.DB.Exec("DELETE FROM Students")
	Config.DB.Exec("DELETE FROM Registers")
}

func populateTestDB() {
	teacher1 := Models.Teacher{Email: "testTeacher1@example.com"}
	teacher2 := Models.Teacher{Email: "testTeacher2@example.com"}
	teacher3 := Models.Teacher{Email: "testTeacher3@example.com"}
	Config.DB.Create(&teacher1)
	Config.DB.Create(&teacher2)
	Config.DB.Create(&teacher3)

	student1 := Models.Student{Email: "testStudent1@example.com", Suspended: false}
	student2 := Models.Student{Email: "testStudent2@example.com", Suspended: false}
	student3 := Models.Student{Email: "testStudent3@example.com", Suspended: false}
	student4 := Models.Student{Email: "testStudent4@example.com", Suspended: false}
	Config.DB.Create(&student1)
	Config.DB.Create(&student2)
	Config.DB.Create(&student3)
	Config.DB.Create(&student4)
}

func TestMain(t *testing.T) {
	setupDB()
	t.Run("Register Functionality", RegisterStudentTest)
	t.Run("Suspend Functionality", SuspendTest)
	t.Run("Common Students Functionality", GetCommonStudentsTest)
	t.Run("Retrieve Receipients Functionality", RetrieveStudentsTest)
	t.Cleanup(clearDB)
}

func RegisterStudentTest(t *testing.T) {
	t.Run("Register a student", RegisterOneStudent)
	t.Run("Register multiple students", RegisterMultipleStudents)
}

func RegisterOneStudent(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/api/register", RegisterStudent)
	jsonParam := `{"teacher":"testTeacher1@example.com","students":["testStudent1@example.com"]}`

	req, err := http.NewRequest(http.MethodPost, "/api/register", strings.NewReader(string(jsonParam)))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code == http.StatusNoContent {
			t.Logf("Expected to get status %d is same as %d\n", http.StatusNoContent, w.Code)
	} else {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusNoContent, w.Code)
	}
}

func RegisterMultipleStudents(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/api/register", RegisterStudent)
	jsonParam := `{"teacher":"testTeacher2@example.com","students":["testStudent1@example.com", "testStudent2@example.com", "testStudent3@example.com"]}`

	req, err := http.NewRequest(http.MethodPost, "/api/register", strings.NewReader(string(jsonParam)))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code == http.StatusNoContent {
			t.Logf("Expected to get status %d is same as %d\n", http.StatusNoContent, w.Code)
	} else {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusNoContent, w.Code)
	}
}

func SuspendTest(t *testing.T) {
	setupDB()
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/api/suspend", SuspendStudent)
	jsonParam := `{"student": "testStudent1@example.com"}`

	req, err := http.NewRequest(http.MethodPost, "/api/suspend", strings.NewReader(string(jsonParam)))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code == http.StatusNoContent {
			t.Logf("Expected to get status %d is same as %d\n", http.StatusNoContent, w.Code)
	} else {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusNoContent, w.Code)
	}
}

func GetCommonStudentsTest(t *testing.T) {
	t.Run("Single Teacher provided", GetSingleCommonStudents)
	t.Run("Multiple Teachers provided", GetMultipleCommonStudents)
}

func GetSingleCommonStudents(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/api/commonstudents", CommonStudents)

	req, err := http.NewRequest(http.MethodGet, "/api/commonstudents?teacher=testTeacher2@example.com", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
			t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
	
	expectedCount := 3
	var data map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	common := data["students"].([]interface{})
	if len(common) == expectedCount {
		t.Logf("Expected to get count %d is same as %d\n", len(common), expectedCount)
	} else {
			t.Fatalf("Expected to get count %d but instead got %d\n", len(common), expectedCount)
	}
}

func GetMultipleCommonStudents(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/api/commonstudents", CommonStudents)

	req, err := http.NewRequest(http.MethodGet, "/api/commonstudents?teacher=testTeacher1@example.com&teacher=testTeacher2@example.com", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
			t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
	
	expectedCount := 1
	var data map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	common := data["students"].([]interface{})
	if len(common) == expectedCount {
		t.Logf("Expected to get count %d is same as %d\n", len(common), expectedCount)
	} else {
			t.Fatalf("Expected to get count %d but instead got %d\n", len(common), expectedCount)
	}
}

func RetrieveStudentsTest(t *testing.T) {
	t.Run("Notifications without tags (ignores suspended students)", RetrieveStudentsNoTags)
	t.Run("Notifications with tags", RetrieveStudentsWithTags)
}

func RetrieveStudentsNoTags(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/api/retrievefornotifications", RetrieveStudentsForNotification)
	jsonParam := `{"teacher": "testTeacher2@example.com", "notification": "Hello!"}`

	req, err := http.NewRequest(http.MethodPost, "/api/retrievefornotifications", strings.NewReader(string(jsonParam)))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
			t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	expectedCount := 2
	var data map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	recipient := data["recipients"].([]interface{})
	if len(recipient) == expectedCount {
		t.Logf("Expected to get count %d is same as %d\n", len(recipient), expectedCount)
	} else {
			t.Fatalf("Expected to get count %d but instead got %d\n", len(recipient), expectedCount)
	}
}

func RetrieveStudentsWithTags(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/api/retrievefornotifications", RetrieveStudentsForNotification)
	jsonParam := `{"teacher": "testTeacher2@example.com", "notification": "Hello! @testStudent4@example.com"}`

	req, err := http.NewRequest(http.MethodPost, "/api/retrievefornotifications", strings.NewReader(string(jsonParam)))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
			t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	expectedCount := 3
	var data map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	recipient := data["recipients"].([]interface{})
	if len(recipient) == expectedCount {
		t.Logf("Expected to get count %d is same as %d\n", len(recipient), expectedCount)
	} else {
			t.Fatalf("Expected to get count %d but instead got %d\n", len(recipient), expectedCount)
	}
}
