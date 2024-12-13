package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mekonger/go-image-generator/internal/controllers"
	"github.com/mekonger/go-image-generator/internal/repo"
	"github.com/mekonger/go-image-generator/internal/services"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRoutes(router *gin.RouterGroup) {
	ur := repo.NewUserRepo()
	us := services.NewUserService(ur)
	controller := controllers.NewUserController(us)

	userRoute := router.Group("/user")
	{
		userRoute.GET("/hello/:name", controller.Hello)
	}
}
