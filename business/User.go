package business

import (
	"gin-gorm-mysql/database"
	"gin-gorm-mysql/models"

	"gorm.io/gorm"
)

var db *gorm.DB = database.InitDb()

func FetchUsers(users *[]models.User) (err error) {
	return db.Find(users).Error
}

func CreateUser(user *models.User) (err error) {
	return db.Create(user).Error
}

func FetchUserById(id int, user *models.User) (err error) {
	return db.Where("id = ?", id).First(user).Error
}

func UpdateUser(id int, user *models.User) (err error) {
	user.ID = id
	return db.Save(user).Error
}

func DeleteUser(id int, user *models.User) (err error) {
	user.ID = id
	return db.Delete(user).Error
}
