package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/YasinOkat/go-school-api/middlewares"
	"github.com/YasinOkat/go-school-api/models"
	"github.com/YasinOkat/go-school-api/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Login godoc
// @Summary Login to get a JWT token
// @Description Login with username and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param login body models.LoginInput true "Login"
// @Success 200 {object} models.TokenResponse
// @Failure 401 {object} models.ErrorResponse "Invalid credentials"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var loginInput models.LoginInput
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	user, err := services.GetUserByUsername(loginInput.Username)

	if err != nil || user == nil || !services.CheckPasswordHash(loginInput.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid credentials"})
		return
	}

	tokenString, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, models.TokenResponse{Token: tokenString})
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.UserCreate true "Create User"
// @Success 201 {object} models.UserCreate
// @Failure 400 {object} models.ErrorResponse
// @Failure 409 {object} models.ErrorResponse "username already exists"
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

// GetUserByID godoc
// @Summary Get a user
// @Description Get a user by ID
// @Tags users
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse "user not found"
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid user ID"})
		return
	}

	user, err := services.GetUserByID(uint(userID))

	if err == services.ErrUserNotFound {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "user not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUsers godoc
// @Summary Fetch all users
// @Description Fetch All users
// @Tags users
// @Produce  json
// @Param active query bool false "Filter by active users"
// @Success 200 {object} []models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/ [get]
func GetUsers(c *gin.Context) {
	active := c.DefaultQuery("active", "false") == "true"
	users, err := services.GetUsers(active)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "some error happened"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func generateJWT(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":      user.ID,
		"user_type_id": user.UserTypeID,
		"exp":          time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(middlewares.JWTSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
