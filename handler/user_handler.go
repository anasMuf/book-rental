package handler

import (
	"book-rental-app/model"
	"book-rental-app/service"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(as service.UserService) *UserHandler {
	return &UserHandler{as}
}

func (h *UserHandler) GetAllUser(c echo.Context) error {
	users, err := h.UserService.GetAllUser()
	if err != nil || len(users) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "No available user"})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByEmail(c echo.Context) error {
	userJWT := c.Get("user").(*jwt.Token)
	claims := userJWT.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	users, err := h.UserService.GetUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "No available user"})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var req model.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	log.Println(req)
	user, err := h.UserService.CreateUser(req)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid input"})
	}
	user, err := h.UserService.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Generate JWT token pakai data user
	jwtSecret := os.Getenv("JWT_SECRET")
	log.Println("JWT_SECRET (user_handler.go):", jwtSecret) // Debug log
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID,
		"email":   user.Email,
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": tokenString,
	})
}
