package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserController is a mock of UserController
type MockUserController struct {
	mock.Mock
}

func (m *MockUserController) FetchUsers(ctx *gin.Context) {
	m.Called(ctx)
}

func (m *MockUserController) CreateUser(ctx *gin.Context) {
	m.Called(ctx)
}

func (m *MockUserController) FetchUserById(ctx *gin.Context) {
	m.Called(ctx)
}

func (m *MockUserController) UpdateUser(ctx *gin.Context) {
	m.Called(ctx)
}

func (m *MockUserController) DeleteUser(ctx *gin.Context) {
	m.Called(ctx)
}

func TestSetupRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUserController := new(MockUserController)
	router := NewRouter(mockUserController)

	engine := gin.New()
	router.SetupRouter(engine)

	mockUserController.On("FetchUsers", mock.Anything).Return()
	mockUserController.On("CreateUser", mock.Anything).Return()
	mockUserController.On("FetchUserById", mock.Anything).Return()
	mockUserController.On("UpdateUser", mock.Anything).Return()
	mockUserController.On("DeleteUser", mock.Anything).Return()

	tests := []struct {
		method   string
		endpoint string
	}{
		{"GET", "/api/v1/users"},
		{"POST", "/api/v1/users"},
		{"GET", "/api/v1/users/1"},
		{"PATCH", "/api/v1/users/1"},
		{"DELETE", "/api/v1/users/1"},
	}

	for _, tt := range tests {
		req, _ := http.NewRequest(tt.method, tt.endpoint, nil)
		w := httptest.NewRecorder()
		engine.ServeHTTP(w, req)

		assert.NotEqual(t, http.StatusNotFound, w.Code, "Endpoint %s %s should exist", tt.method, tt.endpoint)
	}
}
