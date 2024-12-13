package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mekonger/go-image-generator/internal/services"
	"net/http"
)

type UserController interface {
	Hello(c *gin.Context)
}

type userController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userController{UserService: userService}
}

func (uc *userController) Hello(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
	}
	c.JSON(http.StatusOK, uc.UserService.Hello(name))
}
