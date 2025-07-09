package student

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sre-bootcamp/internal/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStudentService struct {
	mock.Mock
}

func (m *MockStudentService) CreateStudent(stu *model.Student) (*model.Student, error) {
	args := m.Called(stu)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Student), args.Error(1)
}

func (m *MockStudentService) GetStudent(id string) (*model.Student, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Student), args.Error(1)
}

func (m *MockStudentService) GetAllStudents() ([]model.Student, error) {
	args := m.Called()
	return args.Get(0).([]model.Student), args.Error(1)
}

func (m *MockStudentService) UpdateStudent(id string, stu *model.Student) (*model.Student, error) {
	args := m.Called(id, stu)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Student), args.Error(1)
}

func (m *MockStudentService) DeleteStudent(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// --- Handler Tests ---

func TestCreateStudentHandler(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(MockStudentService)
	handler := NewHandler(mockService)
	router := gin.New()
	router.POST("/students", handler.CreateStudent)

	inputStudent := model.Student{Name: "Test", Email: "test@test.com"}
	expectedStudent := model.Student{ID: "new-uuid", Name: "Test", Email: "test@test.com"}

	mockInput := model.Student{Name: "Test", Email: "test@test.com"}
	mockService.On("CreateStudent", &mockInput).Return(&expectedStudent, nil).Once()

	// Act
	body, _ := json.Marshal(inputStudent)
	req, _ := http.NewRequest(http.MethodPost, "/students", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)
	var responseStudent model.Student
	err := json.Unmarshal(w.Body.Bytes(), &responseStudent)
	assert.NoError(t, err)
	assert.Equal(t, expectedStudent, responseStudent)
	mockService.AssertExpectations(t)
}

func TestGetStudentHandler_NotFound(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(MockStudentService)
	handler := NewHandler(mockService)
	router := gin.New()
	router.GET("/students/:id", handler.GetStudent)

	mockService.On("GetStudent", "non-existent-id").Return(nil, errors.New("not found")).Once()

	// Act
	req, _ := http.NewRequest(http.MethodGet, "/students/non-existent-id", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "student not found")
	mockService.AssertExpectations(t)
}

func TestGetAllStudentsHandler(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(MockStudentService)
	handler := NewHandler(mockService)
	router := gin.New()
	router.GET("/students", handler.GetAllStudents)

	expectedStudents := []model.Student{
		{ID: "uuid-1", Name: "John Doe", Email: "john@example.com"},
		{ID: "uuid-2", Name: "Jane Doe", Email: "jane@example.com"},
	}
	mockService.On("GetAllStudents").Return(expectedStudents, nil).Once()

	// Act
	req, _ := http.NewRequest(http.MethodGet, "/students", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	var responseStudents []model.Student
	err := json.Unmarshal(w.Body.Bytes(), &responseStudents)
	assert.NoError(t, err)
	assert.Equal(t, expectedStudents, responseStudents)
	mockService.AssertExpectations(t)
}

func TestUpdateStudentHandler(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(MockStudentService)
	handler := NewHandler(mockService)
	router := gin.New()
	router.PUT("/students/:id", handler.UpdateStudent)

	studentID := "existing-uuid"
	updatePayload := model.Student{Name: "John Updated", Email: "john.updated@example.com"}
	expectedStudent := model.Student{ID: studentID, Name: "John Updated", Email: "john.updated@example.com"}

	mockService.On("UpdateStudent", studentID, &updatePayload).Return(&expectedStudent, nil).Once()

	// Act
	body, _ := json.Marshal(updatePayload)
	req, _ := http.NewRequest(http.MethodPut, "/students/"+studentID, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	var responseStudent model.Student
	err := json.Unmarshal(w.Body.Bytes(), &responseStudent)
	assert.NoError(t, err)
	assert.Equal(t, expectedStudent, responseStudent)
	mockService.AssertExpectations(t)
}

func TestDeleteStudentHandler(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(MockStudentService)
	handler := NewHandler(mockService)
	router := gin.New()
	router.DELETE("/students/:id", handler.DeleteStudent)

	studentID := "uuid-to-delete"
	// For a successful deletion, the service should return nil error.
	mockService.On("DeleteStudent", studentID).Return(nil).Once()

	// Act
	req, _ := http.NewRequest(http.MethodDelete, "/students/"+studentID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"deleted": true}`, w.Body.String())
	mockService.AssertExpectations(t)
}
