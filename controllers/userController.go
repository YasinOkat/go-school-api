package controllers

import (
	"net/http"
	"strconv"

	"github.com/YasinOkat/go-school-api/models"
	"github.com/YasinOkat/go-school-api/services"
	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.UserCreate true "Create User"
// @Success 201 {object} models.UserCreate
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.UserCreate
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateUser(user)
	if err == services.ErrUsernameExists {
		c.JSON(http.StatusConflict, models.ErrorResponse{Error: "username already exists"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "user created successfully")
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Produce  json
// @Param id path int true "User ID"
// @Success 204
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse "user not found"
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid user ID"})
		return
	}

	err = services.DeleteUser(uint(userID))

	if err == services.ErrUserNotFound {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "user not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
