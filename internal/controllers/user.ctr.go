package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mekonger/go-image-generator/internal/models"
	"github.com/mekonger/go-image-generator/internal/services"
	"log"
	"net/http"
)

type UserController interface {
	Hello(c *gin.Context)
	PostHello(c *gin.Context)
	UploadFile(c *gin.Context)
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

func (uc *userController) PostHello(c *gin.Context) {
	body := models.Message{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println(body)
	c.JSON(http.StatusAccepted, &body)
}

func (uc *userController) UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	err := c.SaveUploadedFile(file, "/tmp/tempfile")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("Uploaded file successfully: %s", file.Filename))
}
