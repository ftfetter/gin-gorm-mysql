package business

import (
	"gin-gorm-mysql/database"
	"gin-gorm-mysql/models"
)

type UserService interface {
	FetchUsers() (*[]models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	FetchUserById(id int) (*models.User, error)
	UpdateUser(id int, user *models.User) (*models.User, error)
	DeleteUser(id int) error
}

type userServiceImpl struct {
	repository database.UserRepository
}

func NewUserService(repository database.UserRepository) UserService {
	return &userServiceImpl{repository: repository}
}

func (s *userServiceImpl) FetchUsers() (*[]models.User, error) {
	return s.repository.FindAll()
}

func (s *userServiceImpl) CreateUser(user *models.User) (*models.User, error) {
	return s.repository.Create(user)
}

func (s *userServiceImpl) FetchUserById(id int) (*models.User, error) {
	return s.repository.FindByID(id)
}

func (s *userServiceImpl) UpdateUser(id int, user *models.User) (*models.User, error) {
	user.ID = id
	return s.repository.Update(user)
}

func (s *userServiceImpl) DeleteUser(id int) error {
	user := models.User{ID: id}
	_, err := s.repository.Delete(&user)
	return err
}
