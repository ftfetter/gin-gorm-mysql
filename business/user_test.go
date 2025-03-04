package business

import (
	"errors"
	"gin-gorm-mysql/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindAll() (*[]models.User, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).(*[]models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) Create(user *models.User) (*models.User, error) {
	args := m.Called(user)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindByID(id int) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) Update(user *models.User) (*models.User, error) {
	args := m.Called(user)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) Delete(user *models.User) (*models.User, error) {
	args := m.Called(user)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestFetchUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	t.Run("success", func(t *testing.T) {
		mockUsers := []models.User{{ID: 1, Name: "John Doe"}}
		mockRepo.On("FindAll").Return(&mockUsers, nil).Once()

		users, err := service.FetchUsers()
		assert.NoError(t, err)
		assert.Equal(t, &mockUsers, users)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("FindAll").Return(nil, errors.New("error")).Once()

		users, err := service.FetchUsers()
		assert.Error(t, err)
		assert.Nil(t, users)
		mockRepo.AssertExpectations(t)
	})
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	t.Run("success", func(t *testing.T) {
		mockUser := models.User{ID: 1, Name: "John Doe"}
		mockRepo.On("Create", &mockUser).Return(&mockUser, nil).Once()

		createdUser, err := service.CreateUser(&mockUser)
		assert.NoError(t, err)
		assert.Equal(t, &mockUser, createdUser)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUser := models.User{ID: 1, Name: "John Doe"}
		mockRepo.On("Create", &mockUser).Return(nil, errors.New("error")).Once()

		createdUser, err := service.CreateUser(&mockUser)
		assert.Error(t, err)
		assert.Nil(t, createdUser)
		mockRepo.AssertExpectations(t)
	})
}

func TestFetchUserById(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	t.Run("success", func(t *testing.T) {
		mockUser := models.User{ID: 1, Name: "John Doe"}
		mockRepo.On("FindByID", 1).Return(&mockUser, nil).Once()

		user, err := service.FetchUserById(1)
		assert.NoError(t, err)
		assert.Equal(t, &mockUser, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("FindByID", 1).Return(nil, errors.New("error")).Once()

		user, err := service.FetchUserById(1)
		assert.Error(t, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	t.Run("success", func(t *testing.T) {
		mockUser := models.User{ID: 1, Name: "John Doe"}
		mockRepo.On("Update", &mockUser).Return(&mockUser, nil).Once()

		updatedUser, err := service.UpdateUser(1, &mockUser)
		assert.NoError(t, err)
		assert.Equal(t, &mockUser, updatedUser)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUser := models.User{ID: 1, Name: "John Doe"}
		mockRepo.On("Update", &mockUser).Return(nil, errors.New("error")).Once()

		updatedUser, err := service.UpdateUser(1, &mockUser)
		assert.Error(t, err)
		assert.Nil(t, updatedUser)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	t.Run("success", func(t *testing.T) {
		mockUser := models.User{ID: 1}
		mockRepo.On("Delete", &mockUser).Return(&mockUser, nil).Once()

		err := service.DeleteUser(1)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUser := models.User{ID: 1}
		mockRepo.On("Delete", &mockUser).Return(nil, errors.New("error")).Once()

		err := service.DeleteUser(1)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
