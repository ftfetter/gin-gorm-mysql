package controllers

import (
	"gin-gorm-mysql/business"
	"gin-gorm-mysql/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	FetchUsers(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	FetchUserById(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userControllerImpl struct {
	service business.UserService
}

func NewUserController(service business.UserService) UserController {
	return &userControllerImpl{service: service}
}

// FetchUsers godoc
// @Summary      Fetch all users registered
// @Description  Fetch all users registered
// @Tags         users
// @Produce      json
// @Success      200  {object}  models.User
// @Failure      400  {object}  models.HTTPError
// @Failure      404  {object}  models.HTTPError
// @Failure      500  {object}  models.HTTPError
// @Router       /users [get]
func (c *userControllerImpl) FetchUsers(ctx *gin.Context) {
	users, err := c.service.FetchUsers()
	if err != nil {
		models.NewHTTPError(ctx, http.StatusInternalServerError, "It was not possible to fetch all users due to an error", err)
		return
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}

// CreateUser godoc
// @Summary      Register a new user
// @Description  Register a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body      models.User  true  "User to be registered"
// @Success      200  {object}  models.User
// @Failure      400  {object}  models.HTTPError
// @Failure      404  {object}  models.HTTPError
// @Failure      500  {object}  models.HTTPError
// @Router       /users [post]
func (c *userControllerImpl) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		models.NewHTTPError(ctx, http.StatusBadRequest, "Invalid request", err)
		return
	}
	if err := user.ValidateUser(); err != nil {
		models.NewHTTPError(ctx, http.StatusBadRequest, "Invalid request", err)
		return
	}
	newUser, err := c.service.CreateUser(&user)
	if err != nil {
		models.NewHTTPError(ctx, http.StatusInternalServerError, "It was not possible to create the user due to an error", err)
	} else {
		ctx.JSON(http.StatusCreated, newUser)
	}
}

// FetchUserById godoc
// @Summary      Fetch an existing user by its ID
// @Description  Fetch an existing user by its ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int          true  "User ID"
// @Success      200  {object}  models.User
// @Failure      400  {object}  models.HTTPError
// @Failure      404  {object}  models.HTTPError
// @Failure      500  {object}  models.HTTPError
// @Router       /users/{id} [get]
func (c *userControllerImpl) FetchUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		models.NewHTTPError(ctx, http.StatusBadRequest, "Invalid request", err)
		return
	}

	user, err := c.service.FetchUserById(id)
	if err != nil {
		models.NewHTTPError(ctx, http.StatusNotFound, "It was not possible to fetch the user due to an error", err)
		return
	} else {
		ctx.JSON(http.StatusOK, user)
	}
}

// UpdateUser godoc
// @Summary      Update an existing user
// @Description  Update an existing user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int          true  "User ID"
// @Param        user body      models.User  true  "User to be updated"
// @Success      200  {object}  models.User
// @Failure      400  {object}  models.HTTPError
// @Failure      404  {object}  models.HTTPError
// @Failure      500  {object}  models.HTTPError
// @Router       /users/{id} [patch]
func (c *userControllerImpl) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		models.NewHTTPError(ctx, http.StatusBadRequest, "Invalid request.", err)
		return
	}

	var payload models.User
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		models.NewHTTPError(ctx, http.StatusBadRequest, "Invalid request", err)
		return
	}

	user, err := c.service.FetchUserById(id)
	if err != nil {
		models.NewHTTPError(ctx, http.StatusNotFound, "", err)
		return
	}

	user.CopyFrom(payload)
	if err := user.ValidateUser(); err != nil {
		models.NewHTTPError(ctx, http.StatusBadRequest, "", err)
		return
	}

	updatedUser, err := c.service.UpdateUser(id, user)
	if err != nil {
		models.NewHTTPError(ctx, http.StatusInternalServerError, "", err)
		return
	} else {
		ctx.JSON(http.StatusOK, updatedUser)
	}
}

// FetchUserById godoc
// @Summary      Delete an existing user
// @Description  Delete an existing user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int          true  "User ID"
// @Success      200  {object}  models.User
// @Failure      400  {object}  models.HTTPError
// @Failure      404  {object}  models.HTTPError
// @Failure      500  {object}  models.HTTPError
// @Router       /users/{id} [delete]
func (c *userControllerImpl) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		models.NewHTTPError(ctx, http.StatusBadRequest, "", err)
		return
	}

	if err := c.service.DeleteUser(id); err != nil {
		models.NewHTTPError(ctx, http.StatusInternalServerError, "", err)
		return
	} else {
		ctx.JSON(http.StatusNoContent, nil)
	}
}
