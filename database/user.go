package database

import (
	"gin-gorm-mysql/models"

	"gorm.io/gorm"
)

// UserRepository defines the interface for user repository
type UserRepository interface {
	FindAll() (*[]models.User, error)
	Create(user *models.User) (*models.User, error)
	FindByID(id int) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(user *models.User) (*models.User, error)
}

// userRepositoryImpl is the concrete implementation of UserRepository
type userRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	db.AutoMigrate(&models.User{})
	return &userRepositoryImpl{db: db}
}

func (repo *userRepositoryImpl) FindAll() (*[]models.User, error) {
	var usersFound *[]models.User
	if result := repo.db.Find(&usersFound); result.Error != nil {
		return nil, result.Error
	} else {
		return usersFound, nil
	}
}

func (repo *userRepositoryImpl) Create(user *models.User) (*models.User, error) {
	if result := repo.db.Create(&user); result.Error != nil {
		return nil, result.Error
	} else {
		return user, nil
	}
}

func (repo *userRepositoryImpl) FindByID(id int) (*models.User, error) {
	var user *models.User
	if result := repo.db.First(&user, id); result.Error != nil {
		return nil, result.Error
	} else {
		return user, nil
	}
}

func (repo *userRepositoryImpl) Update(user *models.User) (*models.User, error) {
	if result := repo.db.Save(&user); result.Error != nil {
		return nil, result.Error
	} else {
		return user, nil
	}
}

func (repo *userRepositoryImpl) Delete(user *models.User) (*models.User, error) {
	if result := repo.db.Delete(&user); result.Error != nil {
		return nil, result.Error
	} else {
		return user, nil
	}
}
