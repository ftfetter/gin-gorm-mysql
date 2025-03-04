package controllers_test

import (
	"errors"
	"gin-gorm-mysql/controllers"
	"gin-gorm-mysql/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	userJSON        = `{"id": 1, "name": "John Doe", "email": "john@doe.com"}`
	updatedUserJSON = `{"name": "John Doe Updated", "email": "john.updated@doe.com"}`
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) FetchUsers() (*[]models.User, error) {
	args := m.Called()
	if users, ok := args.Get(0).(*[]models.User); ok {
		return users, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserService) CreateUser(user *models.User) (*models.User, error) {
	args := m.Called(user)
	if createdUser, ok := args.Get(0).(*models.User); ok {
		return createdUser, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserService) FetchUserById(id int) (*models.User, error) {
	args := m.Called(id)
	if user, ok := args.Get(0).(*models.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserService) UpdateUser(id int, user *models.User) (*models.User, error) {
	args := m.Called(id, user)
	if updatedUser, ok := args.Get(0).(*models.User); ok {
		return updatedUser, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserService) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestFetchUsers(t *testing.T) {
	mockService := new(MockUserService)
	controller := controllers.NewUserController(mockService)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users", controller.FetchUsers)

	t.Run("success", func(t *testing.T) {
		mockUsers := []models.User{{ID: 1, Name: "John Doe", Email: "john@doe.com"}}
		mockService.On("FetchUsers").Return(&mockUsers, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockService.On("FetchUsers").Return(nil, errors.New("error")).Once()

		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		mockService.AssertExpectations(t)
	})
}

func TestCreateUser(t *testing.T) {
	mockService := new(MockUserService)
	controller := controllers.NewUserController(mockService)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/users", controller.CreateUser)

	t.Run("success", func(t *testing.T) {
		mockUser := models.User{ID: 1, Name: "John Doe", Email: "john@doe.com"}
		mockService.On("CreateUser", &mockUser).Return(&mockUser, nil).Once()

		req, _ := http.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUser := models.User{ID: 1, Name: "John Doe", Email: "john@doe.com"}
		mockService.On("CreateUser", &mockUser).Return(nil, errors.New("error")).Once()

		req, _ := http.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		mockService.AssertExpectations(t)
	})
}

func TestFetchUserById(t *testing.T) {
	mockService := new(MockUserService)
	controller := controllers.NewUserController(mockService)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users/:id", controller.FetchUserById)

	t.Run("success", func(t *testing.T) {
		mockUser := models.User{ID: 1, Name: "John Doe", Email: "john@doe.com"}
		mockService.On("FetchUserById", 1).Return(&mockUser, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockService.On("FetchUserById", 1).Return(nil, errors.New("error")).Once()

		req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	mockService := new(MockUserService)
	controller := controllers.NewUserController(mockService)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PATCH("/users/:id", controller.UpdateUser)

	t.Run("success", func(t *testing.T) {
		mockUser := models.User{ID: 1, Name: "John Doe", Email: "john@doe.com"}
		updatedUser := models.User{ID: 1, Name: "John Doe Updated", Email: "john.updated@doe.com"}
		mockService.On("FetchUserById", 1).Return(&mockUser, nil).Once()
		mockService.On("UpdateUser", 1, &updatedUser).Return(&updatedUser, nil).Once()

		req, _ := http.NewRequest(http.MethodPatch, "/users/1", strings.NewReader(updatedUserJSON))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockService.On("FetchUserById", 1).Return(nil, errors.New("error")).Once()

		req, _ := http.NewRequest(http.MethodPatch, "/users/1", strings.NewReader(updatedUserJSON))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUser := models.User{ID: 1, Name: "John Doe", Email: "john@doe.com"}
		updatedUser := models.User{ID: 1, Name: "John Doe Updated", Email: "john.updated@doe.com"}
		mockService.On("FetchUserById", 1).Return(&mockUser, nil).Once()
		mockService.On("UpdateUser", 1, &updatedUser).Return(nil, errors.New("error")).Once()

		req, _ := http.NewRequest(http.MethodPatch, "/users/1", strings.NewReader(updatedUserJSON))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	mockService := new(MockUserService)
	controller := controllers.NewUserController(mockService)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE("/users/:id", controller.DeleteUser)

	t.Run("success", func(t *testing.T) {
		mockService.On("DeleteUser", 1).Return(nil).Once()

		req, _ := http.NewRequest(http.MethodDelete, "/users/1", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNoContent, resp.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockService.On("DeleteUser", 1).Return(errors.New("error")).Once()

		req, _ := http.NewRequest(http.MethodDelete, "/users/1", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		mockService.AssertExpectations(t)
	})
}
