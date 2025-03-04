package database

import (
	"gin-gorm-mysql/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm database: %v", err)
	}

	return gormDB, mock
}

func TestFindAll(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewUserRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).
		AddRow(1, "John Doe", "john@example.com").
		AddRow(2, "Jane Doe", "jane@example.com")

	mock.ExpectQuery("^SELECT \\* FROM `users`").WillReturnRows(rows)

	users, err := repo.FindAll()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(*users) != 2 {
		t.Errorf("expected 2 users, got %d", len(*users))
	}
}

func TestCreate(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewUserRepository(db)

	user := &models.User{Name: "John Doe", Email: "john@example.com"}

	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `users`").
		WithArgs(user.Name, user.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	createdUser, err := repo.Create(user)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if createdUser.Name != user.Name || createdUser.Email != user.Email {
		t.Errorf("expected user %v, got %v", user, createdUser)
	}
}

func TestFindByID(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewUserRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).
		AddRow(1, "John Doe", "john@example.com")

	mock.ExpectQuery("^SELECT \\* FROM `users` WHERE `users`.`id` = \\?").
		WithArgs(1, 1).
		WillReturnRows(rows)

	user, err := repo.FindByID(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if user.Name != "John Doe" || user.Email != "john@example.com" {
		t.Errorf("expected user John Doe, got %v", user)
	}
}

func TestUpdate(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewUserRepository(db)

	user := &models.User{ID: 1, Name: "John Doe", Email: "john@example.com"}

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE `users` SET `name`=\\?,`email`=\\? WHERE `id` = \\?").
		WithArgs(user.Name, user.Email, user.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	updatedUser, err := repo.Update(user)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if updatedUser.Name != user.Name || updatedUser.Email != user.Email {
		t.Errorf("expected user %v, got %v", user, updatedUser)
	}
}

func TestDelete(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewUserRepository(db)

	user := &models.User{ID: 1, Name: "John Doe", Email: "john@example.com"}

	mock.ExpectBegin()
	mock.ExpectExec("^DELETE FROM `users` WHERE `users`.`id` = \\?").
		WithArgs(user.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	deletedUser, err := repo.Delete(user)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if deletedUser.ID != user.ID {
		t.Errorf("expected user ID %d, got %d", user.ID, deletedUser.ID)
	}
}
